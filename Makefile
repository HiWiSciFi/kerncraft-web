REMOVE = rm -rf

.PHONY: all clean

all:
	@$(MAKE) -s -C frontend all
	@$(MAKE) -s -C backend all

clean:
	@$(MAKE) -s -C frontend clean
	@$(MAKE) -s -C backend clean
	@${REMOVE} out/*
