# Set default number is 0
n?=0

run:
	@cd generator && make && cd ..

generate:
	@./fsg generate -n $(n)

post:
	@./fsg post -n $(n)
