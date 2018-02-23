#! /usr/bin/env bash
export DOCKER_LOCALHOST=$(docker inspect --type container -f '{{.NetworkSettings.Gateway}}' flow)
fn apps config set ocode COMPLETER_BASE_URL http://$DOCKER_LOCALHOST:8081
