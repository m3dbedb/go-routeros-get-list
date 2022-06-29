#!/bin/bash#!/bin/bash

# Router Connection Info
export ROUTER_IP_PORT=your_router_ip:port
export ROUTER_USER=your_router_user
export ROUTER_PWD=your_router_pwd

# Parameters
export PRINT_PARAMETERS="list,address"

go run . --router-ip-port=$ROUTER_IP_PORT \
         --router-user=$ROUTER_USER \
         --router-pwd=$ROUTER_PWD \
         --print-parameters=$PRINT_PARAMETERS
