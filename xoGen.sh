#!/bin/bash

set -e

EXTRA=$1

# SRC to be pwd
SRC=$(realpath $(cd -P "$( dirname "${BASH_SOURCE[0]}" )" && pwd ))

XOBIN=$(which xo)

for i in def_*/config; do
  i=$(dirname $i)

  # skip
  if [ -f $i/skip ]; then
    continue
  fi

  source $i/config

  MODELS=$(expr $i : 'def_\(.*\)$')
  MODELS=xo_${MODELS}
  
  mkdir -p $MODELS
  rm -f $MODELS/*.xo.go

  echo -e "------------------------------------------------------\n$i='$DB'"

  if [ -f $i/pre ]; then
    echo -e "\nsourcing $i/pre"
    source $i/pre
  fi

  #echo -e "\nrm $DB"
  #rm -vf "$DB"

  echo -e "\nsqlite3 $DB < $i/db_schema.sql"
  sqlite3 $EXTRA "$DB" < $i/db_schema.sql

  
  if [ -f $i/db_data.sql ]; then
    (set -ex;
     usql -f $i/db_data.sql $DB
    )
  fi

  echo -e "\nxo schema $DB -o $MODELS"
  $XOBIN $EXTRA schema "$DB" -o $MODELS

  for cq in $i/db_custom-*.sql; do
    cqn=$(expr $cq : '.*/db_custom-\(.*\)\.sql$')
    echo -e "\nxo query $DB -o $MODELS < $cq"
    $XOBIN query $EXTRA \
	   -o $MODELS \
	   -M -B -2 -T $cqn \
	   --type-comment="Custom defined search query $cqn" \
	   "$DB" < $cq
  done

  for cq in $i/db_init-*.sql; do
    echo -e "\nsqlite3 $DB < $cq"
    sqlite3 $EXTRA "$DB" < $cq
  done

done

