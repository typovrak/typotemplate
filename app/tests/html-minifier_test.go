package tests

import (
	"testing"
	"typotemplate/html"
)

func validateHTMLMinifier(t *testing.T, raw string, expected string) {
	res := html.Minifier(raw)

	if res != expected {
		t.Errorf("expected %s (length: %d), got %s (length: %d)", expected, len(expected), res, len(res))
	}
}

func TestHTMLMinifier(t *testing.T) {
	t.Run("minifier_0", func(t *testing.T) {
		raw := "<a></a>"
		expected := "<a></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_1", func(t *testing.T) {
		raw := "< a></a>"
		expected := "<a></a>"

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_2", func(t *testing.T) {
		raw := "< a  ></a>"
		expected := "<a></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_3", func(t *testing.T) {
		raw := "< a>   </a>"
		expected := "<a>   </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_4", func(t *testing.T) {
		raw := "< a>Test a</a>"
		expected := "<a>Test a</a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_5", func(t *testing.T) {
		raw := "< a></  a>"
		expected := "<a></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_6", func(t *testing.T) {
		raw := "< a></a  >"
		expected := "<a></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_7", func(t *testing.T) {
		raw := "< a>  Test  </a>"
		expected := "<a>  Test  </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_8", func(t *testing.T) {
		raw := "<a  href=\" https://mscholz.dev\"></a>"
		expected := "<a href=\"https://mscholz.dev\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_9", func(t *testing.T) {
		raw := "<a  href=\" https://mscholz.dev/blog test coucou \"></a>"
		expected := "<a href=\"https://mscholz.dev/blog test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_10", func(t *testing.T) {
		raw := "<a  title=\"test\"></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_11", func(t *testing.T) {
		raw := "<a  title='test'></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_12", func(t *testing.T) {
		raw := "<a  title='  test '></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_13", func(t *testing.T) {
		raw := "<a  title=test></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_14", func(t *testing.T) {
		raw := "<a  title='test \" '></a>"
		expected := "<a title=\"test &quot;\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_15", func(t *testing.T) {
		raw := "<a  title=test   ></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_16", func(t *testing.T) {
		raw := "<a  title=test ></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_17", func(t *testing.T) {
		raw := "<a  title= test ></a>"
		expected := "<a title=\"test\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_18", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html>
	<head>
		<title>Title of the document</title>
	</head>

	<body>
		The content of the document......
	</body>

</html>
`
		expected := "<!DOCTYPE html><html><head><title>Title of the document</title></head><body>The content of the document......</body></html>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_19", func(t *testing.T) {
		raw := "<a  href=\" https://mscholz.dev/blog    test coucou    \"></a>"
		expected := "<a href=\"https://mscholz.dev/blog    test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_20", func(t *testing.T) {
		raw := "<a>\r</a>"
		expected := "<a></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_21", func(t *testing.T) {
		raw := "<a  href=\" https://mscholz.dev/blog    = test coucou    \"></a>"
		expected := "<a href=\"https://mscholz.dev/blog    = test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_22", func(t *testing.T) {
		raw := "<a  href=\" https://mscholz.dev/blog    =  test coucou    \"></a>"
		expected := "<a href=\"https://mscholz.dev/blog    =  test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_23", func(t *testing.T) {
		raw := `
<  !DOCTYPE		 html  >
<html>
	<head>
		<title>Title of the document</title>
	</head>

	<body>
		The content of the document......
	</body>

</html>
`
		expected := "<!DOCTYPE html><html><head><title>Title of the document</title></head><body>The content of the document......</body></html>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_24", func(t *testing.T) {
		raw := `
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
	<head>
		<title>Title of the document</title>
	</head>

	<body>
		The content of the document......
	</body>

</html>
`
		expected := "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 4.01 Transitional//EN\" \"http://www.w3.org/TR/html4/loose.dtd\"><html><head><title>Title of the document</title></head><body>The content of the document......</body></html>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_25", func(t *testing.T) {
		raw := `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
<html>
	<head>
		<title>Title of the document</title>
	</head>

	<body>
		The content of the document......
	</body>

</html>
`
		expected := "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.1//EN\" \"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd\"><html><head><title>Title of the document</title></head><body>The content of the document......</body></html>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_26", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html>
	<head>
		<title>Title of the document</title>
	</head>

	<body>

		<!-- test -->
		The content of the document......
	</body>

</html>
`
		expected := "<!DOCTYPE html><html><head><title>Title of the document</title></head><body>The content of the document......</body></html>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_27", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html>
	<head>
		<title>Title of the document</title>
	</head>

	<body>

		<p>test <!-- test -->test</p>
		The content of the document......
	</body>

</html>
`
		expected := "<!DOCTYPE html><html><head><title>Title of the document</title></head><body><p>test test</p>The content of the document......</body></html>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_28", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html>
	<head>
		<title>Title of the document</title>
	</head>

	<body>

		<!------------>
		<!------------->
		<!-------------->
		test
		<!-------------->
		<!---->
		<!-- -->
		<p>test <!-- test --> test</p>
		The content of the document......
	</body>

</html>
`
		expected := "<!DOCTYPE html><html><head><title>Title of the document</title></head><body>test<p>test  test</p>The content of the document......</body></html>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_29", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html>
	<head>
		<title>Title of the document</title>
	</head>

	<body>

		<!------------>
		<!------------->
		<!-------------->
		test 
		<!-------------->
		<!---->
		<!-- -->
		<p>test <!-- test --> test</p>
		The content of the document......
	</body>

</html>
`
		expected := "<!DOCTYPE html><html><head><title>Title of the document</title></head><body>test <p>test  test</p>The content of the document......</body></html>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_30", func(t *testing.T) {
		raw := "<!------------>test"
		expected := "test"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_31", func(t *testing.T) {
		raw := "<!------------> test"
		expected := " test"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_32", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<style>
		</style>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<style>
		</style></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_33", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<style>
			.header {
				background: purple;
			}
		</style>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<style>
			.header {
				background: purple;
			}
		</style></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_34", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<  style  type=" text/css " >
			.header {
				background: purple;
			}
		</style>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<style type="text/css">
			.header {
				background: purple;
			}
		</style></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_35", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<style type="text/css">
			.header {
				background: purple;
			}
		</style>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<style type="text/css">
			.header {
				background: purple;
			}
		</style></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_36", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<style media="print">
			.header {
				background: purple;
			}
		</style>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<style media="print">
			.header {
				background: purple;
			}
		</style></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_37", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<style media="screen   "  >
			.header {
				background: purple;
			}
		</  style >
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<style media="screen">
			.header {
				background: purple;
			}
		</style></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_38", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<style media="screen   "  >
			.header {
				background: purple;
			}
		</  style >
		<style media="screen   "  >
			.header {
				background: purple;
			}
		</  style >
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<style media="screen">
			.header {
				background: purple;
			}
		</style><style media="screen">
			.header {
				background: purple;
			}
		</style></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_39", func(t *testing.T) {
		raw := "<a><a> <a></a>"
		expected := "<a><a> <a></a>"
		validateHTMLMinifier(t, raw, expected)
	})
}

//t.Run("title", func(t *testing.T) {
//	t.Errorf("error")
//})
