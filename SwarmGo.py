import commands
import sys
import os
import random
import time

def replace(filename, old, star, var, end):
        s=open(filename).read()
	new = toStr(star), toStr(var), toStr(end)
        if old in s:
                print 'Changed "{old}" to "{new}"'.format(**locals())
                s=s.replace(old, toStr(new))
                f=open(filename, 'w')
                f.write(s)
                f.flush()
                f.close()
        else:
                print '"{old}"Not found. Advise running SwarmReset.py'.format

def log(log, case):
	f = open("LOG/Log.txt", "a")
	f.write('\n')
	f.write(log)
	f.write(': ')
	f.write(case)
	f.close()

def dequote(s):
    if (
        s.startswith(("'", '"')) and s.endswith(("'", '"'))
        and (s[0] == s[-1])  # make sure the pair of quotes match
    ):
        s = s[1:-1]
    return s

def randomChooseType():

	actionchoice = ['"normal",', '"keyed",']
	choice = filter( lambda x: random.randint(0,1), actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('type.go', 'switch choice("normal", "keyed") { //line335R', 'switch choice(',dequote(toStr(choice)),'){ //line335')
	log('type.go line335', toStr(choice))

	actionchoice = ['"normal",', '"keyed",']
	choice = filter( lambda x: random.randint(0,1), actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('type.go', 'switch choice("normal", "keyed") { //line378R', 'switch choice(',dequote(toStr(choice)),'){ //line378')
	log('type.go line378', toStr(choice))

def randomChooseExpr():

	actionchoice = [' "var",', '"indexSlice",', '"indexArray",', '"selector",', '"deref",']
	choice = filter( lambda x: random.randint(0,1) == 0, actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('expr.go', 'switch choice("var", "indexSlice", "indexArray", "selector", "deref") { //line98R', 'switch choice(',dequote(toStr(choice)),') { //line98')
	log('expr.go line98', toStr(choice))

	actionchoice = [' "one",', '"two",', '"slice",']
	choice = filter( lambda x: random.randint(0,1), actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('expr.go', 'switch choice("one", "two", "slice") { //line301R', 'switch choice(',dequote(toStr(choice)),') { //line301')
	log('expr.go line301', toStr(choice))

	actionchoice = [' "int",', '"byteSlice",', '"runeSlice",']
	choice = filter( lambda x: random.randint(0,1), actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('expr.go', 'switch choice("int", "byteSlice", "runeSlice") { //line436R', 'switch choice(',dequote(toStr(choice)),') { //line436')
	log('expr.go line436', toStr(choice))

def randomChooseStmt():

	actionchoice = ['"simple",', '"complex",', '"range",']
	choice = filter( lambda x: random.randint(0,1) == 0, actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('stmt.go', 'switch choice("simple", "complex", "range") {//line92R', 'switch choice(',dequote(toStr(choice)),') { //line92')
	log('stmt.go line92', toStr(choice))

	actionchoice = ['"slice",', '"string",', '"channel",', '"map",']
	choice = filter( lambda x: random.randint(0,1) == 0, actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('stmt.go', 'switch choice("slice", "string", "channel", "map") {//line98R', 'switch choice(',dequote(toStr(choice)),') { //line98')
	log('stmt.go line98', toStr(choice))

	actionchoice = ['"one",', '"two",', '"oneDecl",', '"twoDecl",']
	choice = filter( lambda x: random.randint(0,1) == 0, actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('stmt.go', 'switch choice("one", "two", "oneDecl", "twoDecl") {//line102R', 'switch choice(',dequote(toStr(choice)),') { //line102')
	log('stmt.go line102', toStr(choice))

	actionchoice = ['"one",', '"two",', '"oneDecl",', '"twoDecl",']
	choice = filter( lambda x: random.randint(0,1) == 0, actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('stmt.go', 'switch choice("one", "two", "oneDecl", "twoDecl") {//line121R', 'switch choice(',dequote(toStr(choice)),') { //line121')
	log('stmt.go line121', toStr(choice))

	actionchoice = ['"one",', '"oneDecl",']
	choice = filter( lambda x: random.randint(0,1) == 0, actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('stmt.go', 'switch choice("one", "oneDecl") {//line141R', 'switch choice(',dequote(toStr(choice)),') { //line141')
	log('stmt.go line141', toStr(choice))

	actionchoice = ['"one",', '"two",', '"oneDecl",', '"twoDecl",']
	choice = filter( lambda x: random.randint(0,1) == 0, actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('stmt.go', 'switch choice("one", "two", "oneDecl", "twoDecl") {//line154R', 'switch choice(',dequote(toStr(choice)),') { //line154')
	log('stmt.go line154', toStr(choice))


	actionchoice = ['"empty",', '"inc",', '"assign",', '"oas",', '"send",', '"expr",']
	choice = filter( lambda x: random.randint(0,1) == 0, actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('stmt.go', 'switch choice("empty", "inc", "assign", "oas", "send", "expr") {//line198R', 'switch choice(',dequote(toStr(choice)),') { //line198')
	log('stmt.go line198', toStr(choice))

	actionchoice = ['"normal",', '"decl",']
	choice = filter( lambda x: random.randint(0,1) == 0, actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('stmt.go', 'switch choice("normal", "decl") {//line239R', 'switch choice(',dequote(toStr(choice)),') { //line239')
	log('stmt.go line239', toStr(choice))

	actionchoice = ['"one",', '"two",', '"oneDecl",', '"twoDecl",']
	choice = filter( lambda x: random.randint(0,1) == 0, actionchoice)
	if choice == []:
		choice = random.choice(actionchoice)
	print "Replacing with", choice
	replace('stmt.go', 'switch choice("one", "two", "oneDecl", "twoDecl") {//line306R', 'switch choice(',dequote(toStr(choice)),') { //line306')
	log('stmt.go line306', toStr(choice))

def toStr(x):
  str = ''
  for i in x:
     str += i
  return str


os.system('python reset.py')
print 'Starting'
os.chdir ("..")
os.chdir ("/nfs/guille/groce/users/lamki/Desktop/gosmith/src/gosmith")
os.system ("mkdir LOG")
f = open("LOG/Log.txt", "a")
f.write(time.strftime("%m/%d/%Y")+" "+time.strftime("%H:%M:%S"))
f.close()
randomChooseType()
randomChooseExpr()
randomChooseStmt()
os.system('clear')

print "SwarmGo done. Ready to build Driver.go"
