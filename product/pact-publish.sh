#!/bin/bash

set -x

curl -X PUT \
    http://localhost/pacts/provider/CampaignService/consumer/ProductService/version/${VERSION} \
    -H "Content-Type: application/json" \
    -d @/Users/abdulsamet.ileri/Desktop/personal/cdc-pact-gophercon-2021/product/pacts/productservice-campaignservice.json