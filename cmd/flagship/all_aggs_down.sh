#!/bin/bash
go build .
./flagship throttle set allowQuoteForCTM 0 && ./flagship throttle set allowQuoteForCONFUSED 0 && ./flagship throttle set allowQuoteForGC 0 && ./flagship throttle set allowQuoteForMSM 0
echo "all aggs down"
