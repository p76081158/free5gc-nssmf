#!/bin/bash

if [ -z "$1" ]
then
    echo "Please enter slice id!"
    exit
else
    slice=$1
fi

if [ -z "$2" ]
then
    sfc=false
else
    sfc=$2
fi

if [ "$sfc" = true ]
then
    dir="overlays/sfc"
else
    dir="base"
fi

cd network-slice/$slice
# Delete UPF
kubectl delete -k upf-$slice/$dir/

# Delete SMF
kubectl delete -k smf-$slice/$dir/

# Apply NetworkSlice Custom Resource
kubectl delete -f custom-resource/

# Delete Subnet
kubectl delete -f subnet/

# Delete Network-Attachment-Definition
kubectl delete -f network-attachment-definition/
