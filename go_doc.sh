#!/usr/bin/env bash
( sleep 3;  xdg-open 'http://localhost:6060/pkg/vendingMaxine/?m=all' &>/dev/null & )
godoc -http=localhost:6060 


