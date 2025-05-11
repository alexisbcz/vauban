#!/opt/homebrew/bin/bash
OUTPUT_FILE="tags.go"

cat > "$OUTPUT_FILE" << 'EOL'
// Code generated - DO NOT EDIT.
package tag

import (
	"github.com/alexisbcz/vauban/pkg/html"
)

EOL

declare -A VOID_ELEMENTS
VOID_ELEMENTS[area]=true
VOID_ELEMENTS[base]=true
VOID_ELEMENTS[br]=true
VOID_ELEMENTS[col]=true
VOID_ELEMENTS[embed]=true
VOID_ELEMENTS[hr]=true
VOID_ELEMENTS[img]=true
VOID_ELEMENTS[input]=true
VOID_ELEMENTS[link]=true
VOID_ELEMENTS[meta]=true
VOID_ELEMENTS[param]=true
VOID_ELEMENTS[source]=true
VOID_ELEMENTS[track]=true
VOID_ELEMENTS[wbr]=true

GLOBAL_ATTRIBUTES=(
  "accesskey" "class" "contenteditable" "dir" "draggable"
  "hidden" "id" "lang" "spellcheck" "style" "tabindex" "title"
  "translate" "role"
)

capitalize_first() {
  local input="$1"
  
  special_cases=(svg html)

  for special in "${special_cases[@]}"; do
    if [[ "$input" == "$special" ]]; then
      return
    fi
  done

  
  first=$(echo "${input:0:1}" | tr '[:lower:]' '[:upper:]')
  rest="${input:1}"
  echo "$first$rest"
}

attr_to_method() {
  local attr="$1"
  
  case "$attr" in
    "class") echo "Class" ;;
    "id") echo "ID" ;;
    "src") echo "Src" ;;
    "href") echo "Href" ;;
    *)
      if [[ "$attr" == *"-"* ]]; then
        local result=""
        IFS='-' read -ra PARTS <<< "$attr"
        for i in "${!PARTS[@]}"; do
          if [ $i -eq 0 ]; then
            result+=$(capitalize_first "${PARTS[$i]}")
          else
            result+=$(capitalize_first "${PARTS[$i]}")
          fi
        done
        echo "$result"
      else
        capitalize_first "$attr"
      fi
      ;;
  esac
}

generate_element() {
  local name="$1"
  local struct_name="$1"
  local is_void="${2:-false}"
  local attributes=("${@:3}")

  go_reserved_names=(
  html break default func interface select case defer go map struct
  chan else goto package switch const fallthrough if range type
  continue for import return var

  append bool byte cap close complex complex64 complex128 uint16
  copy false float32 float64 imag int int8 int16 int32 int64
  iota len make new nil panic print println real recover
  rune string true uint uint8 uint32 uint64 uintptr
  )

  for reserved in "${go_reserved_names[@]}"; do
    if [[ "$name" == "$reserved" ]]; then
        struct_name="${name}_"
        break
    fi
  done

  local cap_name=$(capitalize_first "$name")
  
  cat >> "$OUTPUT_FILE" << EOL

type $struct_name struct {
	*Tag
}

func $cap_name(children ...html.Node) *$struct_name {
	return &$struct_name{New("$name", $is_void, children)}
}
EOL

  for attr in "${attributes[@]}"; do
    local method_name=$(attr_to_method "$attr")
    local method_name_if="${method_name}If"
    cat >> "$OUTPUT_FILE" << EOL

func (e *$struct_name) $method_name(value string) *$struct_name {
	e.Attribute("$attr", value)
	return e
}

func (e *$struct_name) $method_name_if(condition bool, value string) *$struct_name {
	if condition {
		e.Attribute("$attr", value)
	}
	return e
}
EOL
  done
}

declare -A SPECIFIC_ATTRIBUTES

SPECIFIC_ATTRIBUTES[a]="href target download rel type hreflang media ping referrerpolicy"

SPECIFIC_ATTRIBUTES[img]="src alt width height loading srcset sizes crossorigin decoding ismap referrerpolicy usemap"

SPECIFIC_ATTRIBUTES[input]="type name value placeholder required disabled readonly checked min max pattern autocomplete autofocus form formaction formenctype formmethod formnovalidate formtarget"

SPECIFIC_ATTRIBUTES[form]="action method enctype target novalidate autocomplete name"

SPECIFIC_ATTRIBUTES[button]="type name value disabled form formaction formmethod formenctype formtarget formnovalidate autofocus"

SPECIFIC_ATTRIBUTES[textarea]="name rows cols placeholder required disabled readonly maxlength minlength wrap autofocus form"

SPECIFIC_ATTRIBUTES[select]="name multiple required disabled size autofocus form"

SPECIFIC_ATTRIBUTES[option]="value selected disabled label"

SPECIFIC_ATTRIBUTES[table]="border cellpadding cellspacing"

SPECIFIC_ATTRIBUTES[th]="scope colspan rowspan headers"

SPECIFIC_ATTRIBUTES[td]="colspan rowspan headers"

SPECIFIC_ATTRIBUTES[iframe]="src width height allow allowfullscreen name sandbox srcdoc loading referrerpolicy"

SPECIFIC_ATTRIBUTES[audio]="src controls autoplay loop muted preload"

SPECIFIC_ATTRIBUTES[video]="src controls autoplay loop muted preload poster width height"

SPECIFIC_ATTRIBUTES[canvas]="width height"

SPECIFIC_ATTRIBUTES[script]="src type async defer crossorigin integrity referrerpolicy nomodule"

SPECIFIC_ATTRIBUTES[link]="href rel type media sizes crossorigin integrity referrerpolicy"

SPECIFIC_ATTRIBUTES[meta]="name content charset http-equiv"

SPECIFIC_ATTRIBUTES[ol]="start type reversed"

SPECIFIC_ATTRIBUTES[li]="value"

SPECIFIC_ATTRIBUTES[blockquote]="cite"

SPECIFIC_ATTRIBUTES[time]="datetime"

SPECIFIC_ATTRIBUTES[progress]="value max"

SPECIFIC_ATTRIBUTES[meter]="value min max low high optimum"

SPECIFIC_ATTRIBUTES[details]="open"

SPECIFIC_ATTRIBUTES[svg]="width height viewBox preserveAspectRatio xmlns version"

SPECIFIC_ATTRIBUTES[path]="d fill stroke stroke-width"

SPECIFIC_ATTRIBUTES[rect]="x y width height rx ry fill stroke"

SPECIFIC_ATTRIBUTES[circle]="cx cy r fill stroke"

SPECIFIC_ATTRIBUTES[line]="x1 y1 x2 y2 stroke stroke-width"

SPECIFIC_ATTRIBUTES[polygon]="points fill stroke"

SPECIFIC_ATTRIBUTES[polyline]="points fill stroke"

SPECIFIC_ATTRIBUTES[text]="x y font-size fill text-anchor"

SPECIFIC_ATTRIBUTES[g]="transform"

SPECIFIC_ATTRIBUTES[label]="for form"

SPECIFIC_ATTRIBUTES[fieldset]="disabled form name"

SPECIFIC_ATTRIBUTES[optgroup]="label disabled"

SPECIFIC_ATTRIBUTES[output]="for form name"

SPECIFIC_ATTRIBUTES[map]="name"

SPECIFIC_ATTRIBUTES[area]="alt coords shape href target download rel"

SPECIFIC_ATTRIBUTES[base]="href target"

SPECIFIC_ATTRIBUTES[col]="span"

SPECIFIC_ATTRIBUTES[colgroup]="span"

HTML_ELEMENTS=(
  "a" "abbr" "address" "area" "article" "aside" "audio" "b" "base" "bdi" "bdo" 
  "blockquote" "body" "br" "button" "canvas" "caption" "cite" "code" "col" 
  "colgroup" "data" "datalist" "dd" "del" "details" "dfn" "dialog" "div" "dl" 
  "dt" "em" "embed" "fieldset" "figcaption" "figure" "footer" "form" "h1" "h2" 
  "h3" "h4" "h5" "h6" "head" "header" "hr" "html" "i" "iframe" "img" "input" 
  "ins" "kbd" "label" "legend" "li" "link" "main" "map" "mark" "meta" "meter" 
  "nav" "noscript" "object" "ol" "optgroup" "option" "output" "p" "param" 
  "picture" "pre" "progress" "q" "rp" "rt" "ruby" "s" "samp" "script" "section" 
  "select" "small" "source" "span" "strong" "style" "sub" "summary" "sup" 
  "svg" "table" "tbody" "td" "template" "textarea" "tfoot" "th" "thead" "time" 
  "title" "tr" "track" "u" "ul" "var" "video" "wbr"
  
  "circle" "ellipse" "g" "line" "path" "polygon" "polyline" "rect" "text" "use"
)

echo "Generating HTML elements code..."

for element in "${HTML_ELEMENTS[@]}"; do
  is_void="${VOID_ELEMENTS[$element]:-false}"
  
  IFS=' ' read -ra specific_attrs <<< "${SPECIFIC_ATTRIBUTES[$element]:-}"
  
  all_attrs=("${specific_attrs[@]}" "${GLOBAL_ATTRIBUTES[@]}")
  
  declare -A attr_map
  unique_attrs=()
  
  for attr in "${all_attrs[@]}"; do
    if [[ -n "$attr" && -z "${attr_map[$attr]}" ]]; then
      attr_map[$attr]=1
      unique_attrs+=("$attr")
    fi
  done
  
  common_attrs=("class" "id" "style" "name" "value")
  for common in "${common_attrs[@]}"; do
    for i in "${!unique_attrs[@]}"; do
      if [[ "${unique_attrs[$i]}" == "$common" ]]; then
        unset "unique_attrs[$i]"
        unique_attrs=("$common" "${unique_attrs[@]}")
        break
      fi
    done
  done
  
  generate_element "$element" "$is_void" "${unique_attrs[@]}"
  
  echo "Generated code for <$element>"
done

echo "Code generation complete! Check $OUTPUT_FILE"
chmod +x "$OUTPUT_FILE"