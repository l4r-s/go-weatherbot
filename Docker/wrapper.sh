#!/bin/bash

if  [ ! -f /data/foo.db ]; then
	echo "---------------- creating sqlite db --------------------"
	mkdir -p /data  
	sqlite3 /data/foo.db << "EOF"
		CREATE TABLE `data` (
		    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
		    `devid` VARCHAR(64) NULL,
		    `temp` DOUBLE NULL,
		    `hum` DOUBLE NULL,
		    `timestamp` DOUBLE NULL
		);
EOF
	/go/bin/go-weatherbot
else
	echo "--------------- using db /data/foo.db -----------------"
	/go/bin/go-weatherbot
fi

