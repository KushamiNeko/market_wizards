css:
	/home/onionhuang/programming_projects/golang/sass_compile/bin/main -input ./templates -output ./resource/css
	/home/onionhuang/programming_projects/golang/web_minifier/bin/main -input ./resource/css -overwrite -nohtml -nojs

js:
	/home/onionhuang/programming_projects/golang/web_minifier/bin/main -input ./templates -output ./resource/js -nohtml -nocss

web:
	make css
	make js
