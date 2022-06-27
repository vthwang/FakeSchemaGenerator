# Set default number is 0
n?=0
f?=1

run:
	@cd generator && make && cd ..

generate:
	@./fsg generate -n $(n)

post:
	@./fsg post -f $(f) -t $(t)
