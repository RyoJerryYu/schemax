package schema

import (
	"github.com/RyoJerryYu/go-utilx/pkg/container/slicex"
	"github.com/RyoJerryYu/go-utilx/pkg/utils/convertx"
	xo "github.com/xo/xo/types"
)

func SetFromXO(s *xo.Set) *Set {
	return &Set{
		Queries: slicex.To(s.Queries, QueryFromXO),
		Schemas: slicex.To(s.Schemas, SchemaFromXO),
	}
}

func QueryFromXO(q xo.Query) *Query {
	return &Query{
		Driver:       q.Driver,
		Name:         q.Name,
		Comment:      q.Comment,
		Exec:         q.Exec,
		Flat:         q.Flat,
		One:          q.One,
		Interpolate:  q.Interpolate,
		Type:         q.Type,
		TypeComment:  q.TypeComment,
		Fields:       slicex.To(q.Fields, FieldFromXO),
		ManualFields: q.ManualFields,
		Params:       slicex.To(q.Params, FieldFromXO),
		Query:        q.Query,
		Comments:     q.Comments,
	}
}

func SchemaFromXO(s xo.Schema) *Schema {
	return &Schema{
		Driver: s.Driver,
		Name:   s.Name,
		Enums:  slicex.To(s.Enums, EnumFromXO),
		Procs:  slicex.To(s.Procs, ProcFromXO),
		Tables: slicex.To(s.Tables, TableFromXO),
		Views:  slicex.To(s.Views, TableFromXO),
	}
}

func EnumFromXO(e xo.Enum) *Enum {
	return &Enum{
		Name:   e.Name,
		Values: slicex.To(e.Values, FieldFromXO),
	}
}

func ProcFromXO(p xo.Proc) *Proc {
	return &Proc{
		Id:         p.ID,
		Type:       p.Type,
		Name:       p.Name,
		Params:     slicex.To(p.Params, FieldFromXO),
		Returns:    slicex.To(p.Returns, FieldFromXO),
		Void:       p.Void,
		Definition: p.Definition,
	}
}

func TableFromXO(t xo.Table) *Table {
	return &Table{
		Type:        t.Type,
		Name:        t.Name,
		Columns:     slicex.To(t.Columns, FieldFromXO),
		PrimaryKeys: slicex.To(t.PrimaryKeys, FieldFromXO),
		Indexes:     slicex.To(t.Indexes, IndexFromXO),
		ForeignKeys: slicex.To(t.ForeignKeys, ForeignKeyFromXO),
		Manual:      t.Manual,
		Definition:  t.Definition,
	}
}

func IndexFromXO(idx xo.Index) *Index {
	return &Index{
		Name:      idx.Name,
		Fields:    slicex.To(idx.Fields, FieldFromXO),
		IsUnique:  idx.IsUnique,
		IsPrimary: idx.IsPrimary,
		Func:      idx.Func,
	}
}

func ForeignKeyFromXO(fk xo.ForeignKey) *ForeignKey {
	return &ForeignKey{
		Name:      fk.Name,
		Fields:    slicex.To(fk.Fields, FieldFromXO),
		RefTable:  fk.RefTable,
		RefFields: slicex.To(fk.RefFields, FieldFromXO),
		Func:      fk.Func,
		RefFunc:   fk.RefFunc,
	}
}

func FieldFromXO(f xo.Field) *Field {
	var constValue *int32
	if f.ConstValue != nil {
		constValue = convertx.Int32Ptr(int32(*f.ConstValue))
	}
	return &Field{
		Name:        f.Name,
		Type:        TypeFromXO(f.Type),
		Default:     f.Default,
		IsPrimary:   f.IsPrimary,
		IsSequence:  f.IsSequence,
		ConstValue:  constValue,
		Interpolate: f.Interpolate,
		Join:        f.Join,
		Comment:     f.Comment,
	}
}

func TypeFromXO(t xo.Type) *Type {
	var enum *Enum
	if t.Enum != nil {
		enum = EnumFromXO(*t.Enum)
	}
	return &Type{
		Type:     t.Type,
		Prec:     int32(t.Prec),
		Scale:    int32(t.Scale),
		Nullable: t.Nullable,
		IsArray:  t.IsArray,
		Unsigned: t.Unsigned,
		Enum:     enum,
	}
}
