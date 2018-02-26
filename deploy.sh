#! /usr/bin/env bash

# deploy all function to local fn server
fn deploy --all --local

# set the COMPLETER_BASE_URL config variable at the application level to
# allow functions using flow to find the flow server
export DOCKER_LOCALHOST=$(docker inspect --type container -f '{{.NetworkSettings.Gateway}}' flow)
fn apps config set ocode COMPLETER_BASE_URL http://$DOCKER_LOCALHOST:8081
