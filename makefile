css:
	/home/onionhuang/programming_projects/golang/sass_compile/bin/main -input ./templates 
	/home/onionhuang/programming_projects/golang/make_template/bin/main -input ./templates
	rm templates/**/*.css

js:
	/home/onionhuang/programming_projects/golang/make_template/bin/main -input ./templates


web:
	make css
