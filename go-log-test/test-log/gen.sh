#!/bin/bash

 loggen --inet --dgram --size 128 --rate $1 --interval 60 127.0.0.1 8514