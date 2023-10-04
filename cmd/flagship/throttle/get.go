package throttle

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/joerdav/flagship/internal/dynamostore"
)

type Get struct {
	Store dynamostore.DynamoStore
	Out   io.Writer
}

func (g Get) Run(args []string) error {
	if len(args) < 1 {
		g.Help()
		return errors.New("No throttleName provided.")
	}
	_, throttles, err := g.Store.Load(context.Background())
	if err != nil {
		return fmt.Errorf("Error loading throttles: %s", err.Error())
	}

	for i := 0; i < len(args); i++ {
		th, ok := throttles[args[i]]
		if !ok {
			return fmt.Errorf("No throttle found: %s", args[i])
		}
		jth, err := json.MarshalIndent(th, "", "\t")
		if err != nil {
			return err
		}

		fmt.Fprintf(g.Out, "\n%s\n", args[i])
		g.Out.Write(jth)
		fmt.Fprintf(g.Out, "\n")
	}
	return nil
}
func (g Get) Help() {
	fmt.Fprintln(g.Out, `usage: flagship throttle get [throttleName]
	Returns the details of a particular throttle.
	Optionally tableName and recordName can be provided. (default=featureFlagStore, features)`)
}
