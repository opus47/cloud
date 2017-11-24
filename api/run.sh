#!/bin/bash

opus47-server \
  --host=0.0.0.0 \
  --tls-port=443 \
  --tls-certificate=cert.pem \
  --tls-key=key.pem

