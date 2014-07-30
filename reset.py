import os
import shutil
import random

os.chdir ("..")
os.chdir ("/nfs/guille/groce/users/lamki/Desktop/gosmith/src/gosmith")


def toStr(x):
  str = ''
  for i in x:
     str += i
  return str

def reset(name):
	os.remove ('/nfs/guille/groce/users/lamki/Desktop/gosmith/src/gosmith/'+name)
	shutil.copy('/nfs/guille/groce/users/lamki/Desktop/gosmith/src/gosmith/backups/'+name, '/nfs/guille/groce/users/lamki/Desktop/gosmith/src/gosmith')
	print name+" has been reset"


reset('context.go')
reset('expr.go')
reset('gosmith.go')
reset('stmt.go')
reset('type.go')

