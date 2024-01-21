#!/bin/bash

echo "---------------"
dayNo=$1

echo "DAY: $dayNo"

if [ $dayNo -lt 1 ]; then
  echo "missing day number"
  read -n 1 -s
  exit 1
fi

mkdir "./day$dayNo"

if [ $? -ne 0 ]; then
  echo "creating ./day$dayNo failed with code $$"
  read -n 1 -s
  exit 1
fi

cat >"./day$dayNo/data.txt" <<EOM
EOM
cat >"./day$dayNo/data_t.txt" <<EOM
EOM

cat >"./day$dayNo/day$dayNo.go" <<EOM
package day$dayNo

import (
	"bufio"
	"os"

	"github.com/unnamedxaer/aoc2023/help"
)

func extractData() (any, any) {

	f, err := os.Open("./day$dayNo/data_t.txt")
	help.IfErr(err)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Bytes()
	}

	help.IfErr(scanner.Err())

	return nil, nil
}

func Ex1() {
	_, _ := extractData()
}

func Ex2() {

}

EOM

echo "Day $dayNo created."


code "./day$dayNo/day$dayNo.go"
code "./day$dayNo/data.txt"
code "./day$dayNo/data_t.txt"