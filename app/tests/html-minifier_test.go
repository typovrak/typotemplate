package tests

import (
	"fmt"
	"testing"
	"typotemplate/html"

	"github.com/typovrak/typotestcolor"
)

func validateHTMLMinifier(t *testing.T, raw string, expected string) {
	res := html.Minifier(raw)
	err := typotestcolor.TestDiff(expected, res, typotestcolor.TestDiffNewDefaultOpts())
	// add raw value
	if err != nil {
		err = fmt.Errorf("%sraw (length: %d): %s", err, len(raw), raw)
	}

	typotestcolor.AssertNoError(t, err)
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

	t.Run("minifier_40", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script>
			console.log("test");
		</script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script>
			console.log("test");
		</script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_41", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<  script  >
			console.log('test');
		</  script >
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script>
			console.log('test');
		</script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_42", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script type="text/javascript">
			console.log("test");
		</script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script type="text/javascript">
			console.log("test");
		</script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_43", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script type="text/javascript  ">
			console.log("test");
		</script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script type="text/javascript">
			console.log("test");
		</script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_44", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module    src="/assets/js/test.js" >
		</script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="/assets/js/test.js"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_45", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script src="  https://code.jquery.com/jquery-3.3.1.slim.min.js"   integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous">
		</script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_46", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="  https://mscholz.dev/blog/bonjour comment ça  va  ? " >
		</script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="https://mscholz.dev/blog/bonjour comment ça  va  ?"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_47", func(t *testing.T) {
		raw := "<a src=\"https://mscholz.dev/blog test coucou\"></a>"
		expected := "<a src=\"https://mscholz.dev/blog test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_48", func(t *testing.T) {
		raw := "<a  src=\" https://mscholz.dev/blog    test coucou    \"></a>"
		expected := "<a src=\"https://mscholz.dev/blog    test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_49", func(t *testing.T) {
		raw := "<a  src=\" https://mscholz.dev/blog    = test coucou    \"></a>"
		expected := "<a src=\"https://mscholz.dev/blog    = test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_50", func(t *testing.T) {
		raw := "<a  src=\" https://mscholz.dev/blog    =  test coucou    \"></a>"
		expected := "<a src=\"https://mscholz.dev/blog    =  test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_51", func(t *testing.T) {
		raw := "<a action=\"https://mscholz.dev/blog test coucou\"></a>"
		expected := "<a action=\"https://mscholz.dev/blog test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_52", func(t *testing.T) {
		raw := "<a  action=\" https://mscholz.dev/blog    test coucou    \"></a>"
		expected := "<a action=\"https://mscholz.dev/blog    test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_53", func(t *testing.T) {
		raw := "<a  action=\" https://mscholz.dev/blog    = test coucou    \"></a>"
		expected := "<a action=\"https://mscholz.dev/blog    = test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_54", func(t *testing.T) {
		raw := "<a  action=\" https://mscholz.dev/blog    =  test coucou    \"></a>"
		expected := "<a action=\"https://mscholz.dev/blog    =  test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_55", func(t *testing.T) {
		raw := "<a data=\"https://mscholz.dev/blog test coucou\"></a>"
		expected := "<a data=\"https://mscholz.dev/blog test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_56", func(t *testing.T) {
		raw := "<a  data=\" https://mscholz.dev/blog    test coucou    \"></a>"
		expected := "<a data=\"https://mscholz.dev/blog    test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_57", func(t *testing.T) {
		raw := "<a  data=\" https://mscholz.dev/blog    = test coucou    \"></a>"
		expected := "<a data=\"https://mscholz.dev/blog    = test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_58", func(t *testing.T) {
		raw := "<a  data=\" https://mscholz.dev/blog    =  test coucou    \"></a>"
		expected := "<a data=\"https://mscholz.dev/blog    =  test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_59", func(t *testing.T) {
		raw := "<script src=\"\"></script>"
		expected := "<script src=\"\"></script>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_60", func(t *testing.T) {
		raw := "<script src=\" \"></script>"
		expected := "<script src=\"\"></script>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_61", func(t *testing.T) {
		raw := "<script src=\"  \"></script>"
		expected := "<script src=\"\"></script>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_62", func(t *testing.T) {
		raw := "<script src=\"	\"></script>"
		expected := "<script src=\"\"></script>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_63", func(t *testing.T) {
		raw := `<script src="
"></script>`
		expected := "<script src=\"\"></script>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_64", func(t *testing.T) {
		raw := "<a  src=\"  https://mscholz.dev/blog    =  test coucou    \"></a>"
		expected := "<a src=\"https://mscholz.dev/blog    =  test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_65", func(t *testing.T) {
		raw := "<a  src='  https://mscholz.dev/blog  \"  =  test coucou    '></a>"
		expected := "<a src=\"https://mscholz.dev/blog  %22  =  test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_66", func(t *testing.T) {
		raw := "<a  src='  https://mscholz.dev/blog  \"test  =  test coucou    '></a>"
		expected := "<a src=\"https://mscholz.dev/blog  %22test  =  test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_67", func(t *testing.T) {
		raw := "<a  src='  https://mscholz.dev/blog  test\"test  =  test coucou    '></a>"
		expected := "<a src=\"https://mscholz.dev/blog  test%22test  =  test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_68", func(t *testing.T) {
		raw := "<a  src='  https://mscholz.dev/blog  test\"  =  test coucou    '></a>"
		expected := "<a src=\"https://mscholz.dev/blog  test%22  =  test coucou\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_69", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="  https://mscholz.dev/blog/bonjour comment ça  va  ? " >
console.log("test");
		</script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="https://mscholz.dev/blog/bonjour comment ça  va  ?"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_70", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="  https://mscholz.dev/blog/bonjour comment ça  va  ? " >
console.log("test");
		</ script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="https://mscholz.dev/blog/bonjour comment ça  va  ?"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_71", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="  https://mscholz.dev/blog/bonjour comment ça  va  ? " >
console.log("test");
		</  script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="https://mscholz.dev/blog/bonjour comment ça  va  ?"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_72", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="  https://mscholz.dev/blog/bonjour comment ça  va  ? " >
console.log("test");
		</ script >
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="https://mscholz.dev/blog/bonjour comment ça  va  ?"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_73", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="  https://mscholz.dev/blog/bonjour comment ça  va  ? " >
console.log("test");
		< / script  >
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="https://mscholz.dev/blog/bonjour comment ça  va  ?"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_74", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="  https://mscholz.dev/blog/bonjour comment ça  va  ? " >
console.log("test");
		<  / script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="https://mscholz.dev/blog/bonjour comment ça  va  ?"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_75", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="  https://mscholz.dev/blog/bonjour comment ça <  va  ? " >
console.log("test");
		<  / script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="https://mscholz.dev/blog/bonjour comment ça <  va  ?"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_76", func(t *testing.T) {
		raw := "<a title=\" < \"></a>"
		expected := "<a title=\"<\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_77", func(t *testing.T) {
		raw := "<a title=\"  <  \"></a>"
		expected := "<a title=\"<\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_78", func(t *testing.T) {
		raw := "<a title=\" t  <  t \"></a>"
		expected := "<a title=\"t < t\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_79", func(t *testing.T) {
		raw := "<a title=\" > \"></a>"
		expected := "<a title=\">\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_80", func(t *testing.T) {
		raw := "<a title=\"  >  \"></a>"
		expected := "<a title=\">\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_81", func(t *testing.T) {
		raw := "<a title=\" t  >  t \"></a>"
		expected := "<a title=\"t > t\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_82", func(t *testing.T) {
		raw := "<a href=\" < \"></a>"
		expected := "<a href=\"<\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_83", func(t *testing.T) {
		raw := "<a href=\"  <  \"></a>"
		expected := "<a href=\"<\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_84", func(t *testing.T) {
		raw := "<a href=\" t  <  t \"></a>"
		expected := "<a href=\"t  <  t\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_85", func(t *testing.T) {
		raw := "<a href=\" > \"></a>"
		expected := "<a href=\">\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_86", func(t *testing.T) {
		raw := "<a href=\"  >  \"></a>"
		expected := "<a href=\">\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_87", func(t *testing.T) {
		raw := "<a href=\" t  >  t \"></a>"
		expected := "<a href=\"t  >  t\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_88", func(t *testing.T) {
		raw := "<a title=<></a>"
		expected := "<a title=\"<\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_89", func(t *testing.T) {
		raw := "<a href=<></a>"
		expected := "<a href=\"<\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_90", func(t *testing.T) {
		raw := "<a title=< ></a>"
		expected := "<a title=\"<\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_91", func(t *testing.T) {
		raw := "<a href=< ></a>"
		expected := "<a href=\"<\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_92", func(t *testing.T) {
		raw := "<a title=<  ></a>"
		expected := "<a title=\"<\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_93", func(t *testing.T) {
		raw := "<a href=<  ></a>"
		expected := "<a href=\"<\"></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_94", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="  https://mscholz.dev/blog/bonjour comment ça  va  ? " >
console.log("</script");
		<  / script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="https://mscholz.dev/blog/bonjour comment ça  va  ?"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_95", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module>
console.log("</script");
		<  / script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module>
console.log("</script");
		</script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_96", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module>
console.log("</sc  r i p   t");
		<  / script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module>
console.log("</sc  r i p   t");
		</script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_97", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module>
console.log("t  e  t  etst");
		<  / script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module>
console.log("t  e  t  etst");
		</script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_98", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="https://mscholz.dev">
console.log("</scaroipat");
		<  / script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="https://mscholz.dev"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_99", func(t *testing.T) {
		raw := "<a title=>></a>"
		expected := "<a title>></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_100", func(t *testing.T) {
		raw := "<a href=>></a>"
		expected := "<a href>></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_101", func(t *testing.T) {
		raw := "<a title=> ></a>"
		expected := "<a title> ></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_102", func(t *testing.T) {
		raw := "<a href=> ></a>"
		expected := "<a href> ></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_103", func(t *testing.T) {
		raw := "<a title=>  ></a>"
		expected := "<a title>  ></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_104", func(t *testing.T) {
		raw := "<a href=>  ></a>"
		expected := "<a href>  ></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_99", func(t *testing.T) {
		raw := "<a title=>  >  </a>"
		expected := "<a title>  >  </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_100", func(t *testing.T) {
		raw := "<a href=>  >  </a>"
		expected := "<a href>  >  </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_99", func(t *testing.T) {
		raw := "<a title=>  > </a>"
		expected := "<a title>  > </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_100", func(t *testing.T) {
		raw := "<a href=>  > </a>"
		expected := "<a href>  > </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_105", func(t *testing.T) {
		raw := "<a href=></a>"
		expected := "<a href></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_106", func(t *testing.T) {
		raw := "<a href= ></a>"
		expected := "<a href></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_107", func(t *testing.T) {
		raw := "<a href=></a>"
		expected := "<a href></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_108", func(t *testing.T) {
		raw := "<a href=  ></a>"
		expected := "<a href></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_109", func(t *testing.T) {
		raw := "<a href></a>"
		expected := "<a href></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_110", func(t *testing.T) {
		raw := "<a href ></a>"
		expected := "<a href></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_111", func(t *testing.T) {
		raw := "<a href  ></a>"
		expected := "<a href></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_112", func(t *testing.T) {
		raw := "<img/>   test"
		expected := "<img>   test"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_113", func(t *testing.T) {
		raw := "<img src/>   test"
		expected := "<img src>   test"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_114", func(t *testing.T) {
		raw := "<img src=/>   test"
		expected := "<img src=\"/\">   test"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_115", func(t *testing.T) {
		raw := "<img src= />   test"
		expected := "<img src=\"/\">   test"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_116", func(t *testing.T) {
		raw := "<img src=  />   test"
		expected := "<img src=\"/\">   test"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_117", func(t *testing.T) {
		raw := "< img src=  />   test"
		expected := "<img src=\"/\">   test"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_118", func(t *testing.T) {
		raw := "<  img src=  />   test"
		expected := "<img src=\"/\">   test"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_119", func(t *testing.T) {
		raw := "<img src=  / >   test"
		expected := "<img src=\"/\">   test"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_120", func(t *testing.T) {
		raw := "<img src=  /  >   test"
		expected := "<img src=\"/\">   test"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_121", func(t *testing.T) {
		raw := "<a><</a>"
		expected := "<a><</a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_122", func(t *testing.T) {
		raw := "<a> <</a>"
		expected := "<a> <</a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_123", func(t *testing.T) {
		raw := "<a>  <</a>"
		expected := "<a>  <</a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_124", func(t *testing.T) {
		raw := "<a>< </a>"
		expected := "<a>< </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_125", func(t *testing.T) {
		raw := "<a><  </a>"
		expected := "<a><  </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_126", func(t *testing.T) {
		raw := "<a> < </a>"
		expected := "<a> < </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_127", func(t *testing.T) {
		raw := "<a>  <  </a>"
		expected := "<a>  <  </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_128", func(t *testing.T) {
		raw := "<a>></a>"
		expected := "<a>></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_129", func(t *testing.T) {
		raw := "<a> ></a>"
		expected := "<a> ></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_130", func(t *testing.T) {
		raw := "<a>  ></a>"
		expected := "<a>  ></a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_131", func(t *testing.T) {
		raw := "<a>> </a>"
		expected := "<a>> </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_132", func(t *testing.T) {
		raw := "<a>>  </a>"
		expected := "<a>>  </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_133", func(t *testing.T) {
		raw := "<a> > </a>"
		expected := "<a> > </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_134", func(t *testing.T) {
		raw := "<a>  >  </a>"
		expected := "<a>  >  </a>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_135", func(t *testing.T) {
		raw := `<style>
	.header {
		content: "</style";
	}
</style>`

		expected := `<style>
	.header {
		content: "</style";
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_136", func(t *testing.T) {
		raw := `<style>
	.header {
		content: "<\/style>";
	}
</style>`

		expected := `<style>
	.header {
		content: "<\/style>";
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_137", func(t *testing.T) {
		raw := `<style>
	.header {
		content: "<\/style>\"";
	}
</style>`

		expected := `<style>
	.header {
		content: "<\/style>\"";
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_138", func(t *testing.T) {
		raw := `<style>
	.header {
		content: '</style';
	}
</style>`

		expected := `<style>
	.header {
		content: '</style';
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_139", func(t *testing.T) {
		raw := `<style>
	.header {
		content: '<\/style>\'';
	}
</style>`

		expected := `<style>
	.header {
		content: '<\/style>\'';
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_140", func(t *testing.T) {
		raw := `<style>
	.header {
		content: '<\/style>\"';
	}
</style>`

		expected := `<style>
	.header {
		content: '<\/style>\"';
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_141", func(t *testing.T) {
		raw := `<style>
	.header {
		content: "<\/style>\'";
	}
</style>`

		expected := `<style>
	.header {
		content: "<\/style>\'";
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_142", func(t *testing.T) {
		raw := `<style>
	.header {
		content: '<\/style>"';
	}
</style>`

		expected := `<style>
	.header {
		content: '<\/style>"';
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_143", func(t *testing.T) {
		raw := `<style>
	.header {
		content: "<\/style>'";
	}
</style>`

		expected := `<style>
	.header {
		content: "<\/style>'";
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_144", func(t *testing.T) {
		raw := `<style>
	.header {
		content: "t  e   s t";
	}
</style>`

		expected := `<style>
	.header {
		content: "t  e   s t";
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_145", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="/test.js">
console.log("</sc  r i p   t");
		<  / script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="/test.js"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_146", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="/test.js">
console.log("</script");
		</script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="/test.js"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_147", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="/test.js">
console.log("<\/sc  r i p   t>");
		<  / script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="/test.js"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_148", func(t *testing.T) {
		raw := `
<!DOCTYPE html>
<html lang="fr">
	<head>
		<title>Title of the document</title>
	</head>
	<body>
		The content of the document......
		<script module src="/test.js">
console.log("<\/script>");
		</script>
	</body>
</html>
`
		expected := `<!DOCTYPE html><html lang="fr"><head><title>Title of the document</title></head><body>The content of the document......<script module src="/test.js"></script></body></html>`
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_149", func(t *testing.T) {
		raw := `<style>
	.header {
		content: "</s ty   le";
	}
</style>`

		expected := `<style>
	.header {
		content: "</s ty   le";
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_150", func(t *testing.T) {
		raw := `<style>
	.header {
		content: "</s ty   le";
	}

	.footer {
		content: "</s ty   le";
	}
</style>`

		expected := `<style>
	.header {
		content: "</s ty   le";
	}

	.footer {
		content: "</s ty   le";
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_151", func(t *testing.T) {
		raw := `<style>
	.header {
		content: "</style";
	}

	.footer {
		content: "</style";
	}
</style>`

		expected := `<style>
	.header {
		content: "</style";
	}

	.footer {
		content: "</style";
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_152", func(t *testing.T) {
		raw := `<style>
	.header {
		content: "</s ty   le>";
	}

	.footer {
		content: "</s ty   le>";
	}
</style>`

		expected := `<style>
	.header {
		content: "</s ty   le>";
	}

	.footer {
		content: "</s ty   le>";
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_153", func(t *testing.T) {
		raw := `<style>
	.header {
		content: "<\/style>";
	}

	.footer {
		content: "<\/style>";
	}
</style>`

		expected := `<style>
	.header {
		content: "<\/style>";
	}

	.footer {
		content: "<\/style>";
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_154", func(t *testing.T) {
		raw := `<style>
	.header {
		content: "</ s ty   le>";
	}

	.footer {
		content: "</ s ty   le>";
	}
</style>`

		expected := `<style>
	.header {
		content: "</ s ty   le>";
	}

	.footer {
		content: "</ s ty   le>";
	}
</style>`

		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_155", func(t *testing.T) {
		raw := "<img/>"
		expected := "<img>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_156", func(t *testing.T) {
		raw := "< img/>"
		expected := "<img>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_157", func(t *testing.T) {
		raw := "<  img/>"
		expected := "<img>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_158", func(t *testing.T) {
		raw := "<  img / >"
		expected := "<img>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_159", func(t *testing.T) {
		raw := "<  img  /  >"
		expected := "<img>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_160", func(t *testing.T) {
		raw := "<div/>"
		expected := "<div>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_161", func(t *testing.T) {
		raw := "< div/>"
		expected := "<div>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_162", func(t *testing.T) {
		raw := "<  div/>"
		expected := "<div>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_163", func(t *testing.T) {
		raw := "<  div / >"
		expected := "<div>"
		validateHTMLMinifier(t, raw, expected)
	})

	t.Run("minifier_164", func(t *testing.T) {
		raw := "<  div  /  >"
		expected := "<div>"
		validateHTMLMinifier(t, raw, expected)
	})

	// TODO: <a href=<></a>
	// TODO: <img title= test / >
}
