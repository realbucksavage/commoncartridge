package commoncartridge

//go:generate echo "Generating Manifest, Item, Resource, ..."
//go:generate bash -c "goxmlstruct -use-pointers-for-optional-fields=false -named-root -named-types -top-level-attributes -char-data-field-name Text -package-name types -output ./types/structs.go ./types/examples/*.xml"

//go:generate echo "...done!"
