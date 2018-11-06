css:
	~/programming_projects/golang/css_compile/bin/main -input ./templates 
	~/programming_projects/golang/make_template/bin/main -input ./templates
	rm templates/**/*.css

js:
	~/programming_projects/golang/make_template/bin/main -input ./templates


web:
	make css
