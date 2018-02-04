import commands
import sys
import os


os.system ("go build gosmith")
print "Gosmith built"

os.system ("go run driver.go -checkers=amd64,386,arm,exec")
print "Running Test"
