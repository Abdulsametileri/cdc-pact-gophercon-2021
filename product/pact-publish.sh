#!/bin/bash

set -x

VERSION=$1 #like 1.0.0

curl -X PUT \
    http://localhost/pacts/provider/CampaignService/consumer/ProductService/version/${VERSION} \
    -H "Content-Type: application/json" \
    -d @/Users/abdulsamet.ileri/Desktop/personal/cdc-pact-gophercon-2021/product/pacts/productservice-campaignservice.json