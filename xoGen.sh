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
  MODELS=xo${MODELS^}
  
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

  echo -e "\nxo $DB -o $MODELS"
  $XOBIN $EXTRA "$DB" -o $MODELS

  for cq in $i/db_custom-*.sql; do
    cqn=$(expr $cq : '.*/db_custom-\(.*\)\.sql$')
    echo -e "\nxo $DB -o $MODELS < $cq"
    $XOBIN $EXTRA \
	   -o $MODELS \
	   -N -M -B -T $cqn \
	   --query-type-comment="Custom defined search query $cqn" \
	   "$DB" < $cq
  done

  for cq in $i/db_init-*.sql; do
    echo -e "\nsqlite3 $DB < $cq"
    sqlite3 $EXTRA "$DB" < $cq
  done

done

