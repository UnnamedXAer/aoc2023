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
package day${dayNo}

func Ex1() {

}

func Ex2() {

}
EOM

echo "Day $dayNo created."


code "./day$dayNo/day$dayNo.go"
code "./day$dayNo/data.txt"
code "./day$dayNo/data_t.txt"