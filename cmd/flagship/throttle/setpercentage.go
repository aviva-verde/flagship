package throttle

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/joerdav/flagship/internal/dynamostore"
)

type SetPercentage struct {
	Store dynamostore.DynamoStore
	Out   io.Writer
}

func (s SetPercentage) Run(args []string) error {
	if args[0] == "" {
		s.Help()
		return errors.New("No throttleName provided.")
	}
	if args[1] == "" {
		s.Help()
		return errors.New("No probability provided.")
	}
	err := s.Store.SetThrottleProbability(context.Background(), args[0], args[1])
	if err != nil {
		return fmt.Errorf("Error when setting throttle probability: %s", err.Error())
	}
	_, th, err := s.Store.Load(context.Background())
	if err != nil {
		return fmt.Errorf("Error loading throttles: %s", err.Error())
	}
	jth, err := json.MarshalIndent(th[args[0]], "", "\t")
	if err != nil {
		return err
	}

	fmt.Fprintf(s.Out, "\n%s\n", args[0])
	s.Out.Write(jth)
	fmt.Fprintf(s.Out, "\n")
	return nil
}
func (s SetPercentage) Help() {
	fmt.Println(`usage: flagship throttle set [throttleName] [probability]
	Sets a percentage of a throttle.`)
}
