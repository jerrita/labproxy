#!/bin/bash

systemctl daemon-reload
systemctl enable labproxy.service
systemctl start labproxy.service

