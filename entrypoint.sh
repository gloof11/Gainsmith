#!/bin/sh
service nginx start
/pb/pocketbase serve --http=0.0.0.0:8090
