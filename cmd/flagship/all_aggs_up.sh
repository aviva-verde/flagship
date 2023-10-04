#!/bin/bash
go build .
./flagship throttle set allowQuoteForCTM 100 && ./flagship throttle set allowQuoteForCONFUSED 100 && ./flagship throttle set allowQuoteForGC 100 && ./flagship throttle set allowQuoteForMSM 100
echo "all aggs up"
