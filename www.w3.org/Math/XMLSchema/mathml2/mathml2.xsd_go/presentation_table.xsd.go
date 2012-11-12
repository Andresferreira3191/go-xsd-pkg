//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.w3.org/Math/XMLSchema/mathml2/presentation/table.xsd
package gopkg_WwwW3OrgMathXMLSchemaMathml2Mathml2Xsd

//	This is an XML Schema module for tables in MathML presentation.
//	Author: Stéphane Dalmas, INRIA.
import (
	xsdt "github.com/metaleap/go-xsd/types"
)

type TxsdTableAlignmentAttribRowalign xsdt.String

//	Since TxsdTableAlignmentAttribRowalign is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdTableAlignmentAttribRowalign) SetFromString (s string) { (*xsdt.String)(me).SetFromString(s) }

//	Since TxsdTableAlignmentAttribRowalign is just a simple String type, this merely returns the current string value.
func (me TxsdTableAlignmentAttribRowalign) String () string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdTableAlignmentAttribRowalign's alias type xsdt.String
func (me TxsdTableAlignmentAttribRowalign) ToXsdtString () xsdt.String { return xsdt.String(me) }

type XsdGoPkgHasAttr_Rowalign_TxsdTableAlignmentAttribRowalign_Baseline struct {
	Rowalign TxsdTableAlignmentAttribRowalign `xml:"http://www.w3.org/1998/Math/MathML rowalign,attr"`
}

//	Returns the Default value for Rowalign -- "baseline"
func (me *XsdGoPkgHasAttr_Rowalign_TxsdTableAlignmentAttribRowalign_Baseline) RowalignDefault () TxsdTableAlignmentAttribRowalign { return TxsdTableAlignmentAttribRowalign("baseline") }

type TxsdTableAlignmentAttribColumnalign xsdt.String

//	Since TxsdTableAlignmentAttribColumnalign is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdTableAlignmentAttribColumnalign) SetFromString (s string) { (*xsdt.String)(me).SetFromString(s) }

//	Since TxsdTableAlignmentAttribColumnalign is just a simple String type, this merely returns the current string value.
func (me TxsdTableAlignmentAttribColumnalign) String () string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdTableAlignmentAttribColumnalign's alias type xsdt.String
func (me TxsdTableAlignmentAttribColumnalign) ToXsdtString () xsdt.String { return xsdt.String(me) }

type XsdGoPkgHasAttr_Columnalign_TxsdTableAlignmentAttribColumnalign_Center struct {
	Columnalign TxsdTableAlignmentAttribColumnalign `xml:"http://www.w3.org/1998/Math/MathML columnalign,attr"`
}

//	Returns the Default value for Columnalign -- "center"
func (me *XsdGoPkgHasAttr_Columnalign_TxsdTableAlignmentAttribColumnalign_Center) ColumnalignDefault () TxsdTableAlignmentAttribColumnalign { return TxsdTableAlignmentAttribColumnalign("center") }

type XsdGoPkgHasAttr_Groupalign_XsdtString_ struct {
	Groupalign xsdt.String `xml:"http://www.w3.org/1998/Math/MathML groupalign,attr"`
}

type XsdGoPkgHasAtts_TableAlignmentAttrib struct {
	XsdGoPkgHasAttr_Rowalign_TxsdTableAlignmentAttribRowalign_Baseline
	XsdGoPkgHasAttr_Columnalign_TxsdTableAlignmentAttribColumnalign_Center
	XsdGoPkgHasAttr_Groupalign_XsdtString_
}

type XsdGoPkgHasAtts_MtrAttlist struct {
	XsdGoPkgHasAtts_TableAlignmentAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_MlabeledtrAttlist struct {
	XsdGoPkgHasAtts_TableAlignmentAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAttr_Columnspan_XsdtPositiveInteger_1 struct {
	Columnspan xsdt.PositiveInteger `xml:"http://www.w3.org/1998/Math/MathML columnspan,attr"`
}

//	Returns the Default value for Columnspan -- 1
func (me *XsdGoPkgHasAttr_Columnspan_XsdtPositiveInteger_1) ColumnspanDefault () xsdt.PositiveInteger { return xsdt.PositiveInteger(1) }

type XsdGoPkgHasAttr_Rowspan_XsdtPositiveInteger_1 struct {
	Rowspan xsdt.PositiveInteger `xml:"http://www.w3.org/1998/Math/MathML rowspan,attr"`
}

//	Returns the Default value for Rowspan -- 1
func (me *XsdGoPkgHasAttr_Rowspan_XsdtPositiveInteger_1) RowspanDefault () xsdt.PositiveInteger { return xsdt.PositiveInteger(1) }

type XsdGoPkgHasAtts_MtdAttlist struct {
	XsdGoPkgHasAtts_TableAlignmentAttrib

	XsdGoPkgHasAtts_CommonAttrib

	XsdGoPkgHasAttr_Columnspan_XsdtPositiveInteger_1
	XsdGoPkgHasAttr_Rowspan_XsdtPositiveInteger_1
}

type XsdGoPkgHasAttr_Align_XsdtString_Axis struct {
	Align xsdt.String `xml:"http://www.w3.org/1998/Math/MathML align,attr"`
}

//	Returns the Default value for Align -- "axis"
func (me *XsdGoPkgHasAttr_Align_XsdtString_Axis) AlignDefault () xsdt.String { return xsdt.String("axis") }

type TxsdMtableAttlistAlignmentscope xsdt.String

//	Since TxsdMtableAttlistAlignmentscope is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdMtableAttlistAlignmentscope) SetFromString (s string) { (*xsdt.String)(me).SetFromString(s) }

//	Since TxsdMtableAttlistAlignmentscope is just a simple String type, this merely returns the current string value.
func (me TxsdMtableAttlistAlignmentscope) String () string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdMtableAttlistAlignmentscope's alias type xsdt.String
func (me TxsdMtableAttlistAlignmentscope) ToXsdtString () xsdt.String { return xsdt.String(me) }

type XsdGoPkgHasAttr_Alignmentscope_TxsdMtableAttlistAlignmentscope_True struct {
	Alignmentscope TxsdMtableAttlistAlignmentscope `xml:"http://www.w3.org/1998/Math/MathML alignmentscope,attr"`
}

//	Returns the Default value for Alignmentscope -- "true"
func (me *XsdGoPkgHasAttr_Alignmentscope_TxsdMtableAttlistAlignmentscope_True) AlignmentscopeDefault () TxsdMtableAttlistAlignmentscope { return TxsdMtableAttlistAlignmentscope("true") }

type XsdGoPkgHasAttr_Columnwidth_XsdtString_Auto struct {
	Columnwidth xsdt.String `xml:"http://www.w3.org/1998/Math/MathML columnwidth,attr"`
}

//	Returns the Default value for Columnwidth -- "auto"
func (me *XsdGoPkgHasAttr_Columnwidth_XsdtString_Auto) ColumnwidthDefault () xsdt.String { return xsdt.String("auto") }

type XsdGoPkgHasAttr_Width_XsdtString_Auto struct {
	Width xsdt.String `xml:"http://www.w3.org/1998/Math/MathML width,attr"`
}

//	Returns the Default value for Width -- "auto"
func (me *XsdGoPkgHasAttr_Width_XsdtString_Auto) WidthDefault () xsdt.String { return xsdt.String("auto") }

type XsdGoPkgHasAttr_Rowspacing_XsdtString_10Ex struct {
	Rowspacing xsdt.String `xml:"http://www.w3.org/1998/Math/MathML rowspacing,attr"`
}

//	Returns the Default value for Rowspacing -- "1.0ex"
func (me *XsdGoPkgHasAttr_Rowspacing_XsdtString_10Ex) RowspacingDefault () xsdt.String { return xsdt.String("1.0ex") }

type XsdGoPkgHasAttr_Columnspacing_XsdtString_08Em struct {
	Columnspacing xsdt.String `xml:"http://www.w3.org/1998/Math/MathML columnspacing,attr"`
}

//	Returns the Default value for Columnspacing -- "0.8em"
func (me *XsdGoPkgHasAttr_Columnspacing_XsdtString_08Em) ColumnspacingDefault () xsdt.String { return xsdt.String("0.8em") }

type XsdGoPkgHasAttr_Rowlines_XsdtString_None struct {
	Rowlines xsdt.String `xml:"http://www.w3.org/1998/Math/MathML rowlines,attr"`
}

//	Returns the Default value for Rowlines -- "none"
func (me *XsdGoPkgHasAttr_Rowlines_XsdtString_None) RowlinesDefault () xsdt.String { return xsdt.String("none") }

type XsdGoPkgHasAttr_Columnlines_XsdtString_None struct {
	Columnlines xsdt.String `xml:"http://www.w3.org/1998/Math/MathML columnlines,attr"`
}

//	Returns the Default value for Columnlines -- "none"
func (me *XsdGoPkgHasAttr_Columnlines_XsdtString_None) ColumnlinesDefault () xsdt.String { return xsdt.String("none") }

type TxsdMtableAttlistFrame xsdt.String

//	Since TxsdMtableAttlistFrame is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdMtableAttlistFrame) SetFromString (s string) { (*xsdt.String)(me).SetFromString(s) }

//	Since TxsdMtableAttlistFrame is just a simple String type, this merely returns the current string value.
func (me TxsdMtableAttlistFrame) String () string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdMtableAttlistFrame's alias type xsdt.String
func (me TxsdMtableAttlistFrame) ToXsdtString () xsdt.String { return xsdt.String(me) }

//	Returns true if the value of this enumerated TxsdMtableAttlistFrame is "none".
func (me TxsdMtableAttlistFrame) IsNone () bool { return me == "none" }

//	Returns true if the value of this enumerated TxsdMtableAttlistFrame is "solid".
func (me TxsdMtableAttlistFrame) IsSolid () bool { return me == "solid" }

//	Returns true if the value of this enumerated TxsdMtableAttlistFrame is "dashed".
func (me TxsdMtableAttlistFrame) IsDashed () bool { return me == "dashed" }

type XsdGoPkgHasAttr_Frame_TxsdMtableAttlistFrame_None struct {
	Frame TxsdMtableAttlistFrame `xml:"http://www.w3.org/1998/Math/MathML frame,attr"`
}

//	Returns the Default value for Frame -- "none"
func (me *XsdGoPkgHasAttr_Frame_TxsdMtableAttlistFrame_None) FrameDefault () TxsdMtableAttlistFrame { return TxsdMtableAttlistFrame("none") }

type XsdGoPkgHasAttr_Framespacing_XsdtString_04Em05Ex struct {
	Framespacing xsdt.String `xml:"http://www.w3.org/1998/Math/MathML framespacing,attr"`
}

//	Returns the Default value for Framespacing -- "0.4em 0.5ex"
func (me *XsdGoPkgHasAttr_Framespacing_XsdtString_04Em05Ex) FramespacingDefault () xsdt.String { return xsdt.String("0.4em 0.5ex") }

type XsdGoPkgHasAttr_Equalrows_XsdtBoolean_False struct {
	Equalrows xsdt.Boolean `xml:"http://www.w3.org/1998/Math/MathML equalrows,attr"`
}

//	Returns the Default value for Equalrows -- false
func (me *XsdGoPkgHasAttr_Equalrows_XsdtBoolean_False) EqualrowsDefault () xsdt.Boolean { return xsdt.Boolean(false) }

type XsdGoPkgHasAttr_Equalcolumns_XsdtBoolean_False struct {
	Equalcolumns xsdt.Boolean `xml:"http://www.w3.org/1998/Math/MathML equalcolumns,attr"`
}

//	Returns the Default value for Equalcolumns -- false
func (me *XsdGoPkgHasAttr_Equalcolumns_XsdtBoolean_False) EqualcolumnsDefault () xsdt.Boolean { return xsdt.Boolean(false) }

type XsdGoPkgHasAttr_Displaystyle_XsdtBoolean_False struct {
	Displaystyle xsdt.Boolean `xml:"http://www.w3.org/1998/Math/MathML displaystyle,attr"`
}

//	Returns the Default value for Displaystyle -- false
func (me *XsdGoPkgHasAttr_Displaystyle_XsdtBoolean_False) DisplaystyleDefault () xsdt.Boolean { return xsdt.Boolean(false) }

type TxsdMtableAttlistSide xsdt.String

//	Since TxsdMtableAttlistSide is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdMtableAttlistSide) SetFromString (s string) { (*xsdt.String)(me).SetFromString(s) }

//	Since TxsdMtableAttlistSide is just a simple String type, this merely returns the current string value.
func (me TxsdMtableAttlistSide) String () string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdMtableAttlistSide's alias type xsdt.String
func (me TxsdMtableAttlistSide) ToXsdtString () xsdt.String { return xsdt.String(me) }

//	Returns true if the value of this enumerated TxsdMtableAttlistSide is "left".
func (me TxsdMtableAttlistSide) IsLeft () bool { return me == "left" }

//	Returns true if the value of this enumerated TxsdMtableAttlistSide is "right".
func (me TxsdMtableAttlistSide) IsRight () bool { return me == "right" }

//	Returns true if the value of this enumerated TxsdMtableAttlistSide is "leftoverlap".
func (me TxsdMtableAttlistSide) IsLeftoverlap () bool { return me == "leftoverlap" }

//	Returns true if the value of this enumerated TxsdMtableAttlistSide is "rightoverlap".
func (me TxsdMtableAttlistSide) IsRightoverlap () bool { return me == "rightoverlap" }

type XsdGoPkgHasAttr_Side_TxsdMtableAttlistSide_Right struct {
	Side TxsdMtableAttlistSide `xml:"http://www.w3.org/1998/Math/MathML side,attr"`
}

//	Returns the Default value for Side -- "right"
func (me *XsdGoPkgHasAttr_Side_TxsdMtableAttlistSide_Right) SideDefault () TxsdMtableAttlistSide { return TxsdMtableAttlistSide("right") }

type XsdGoPkgHasAttr_Minlabelspacing_TlengthWithUnit_08Em struct {
	Minlabelspacing TlengthWithUnit `xml:"http://www.w3.org/1998/Math/MathML minlabelspacing,attr"`
}

//	Returns the Default value for Minlabelspacing -- "0.8em"
func (me *XsdGoPkgHasAttr_Minlabelspacing_TlengthWithUnit_08Em) MinlabelspacingDefault () TlengthWithUnit { return TlengthWithUnit("0.8em") }

type XsdGoPkgHasAtts_MtableAttlist struct {
	XsdGoPkgHasAtts_TableAlignmentAttrib

	XsdGoPkgHasAtts_CommonAttrib

	XsdGoPkgHasAttr_Align_XsdtString_Axis
	XsdGoPkgHasAttr_Alignmentscope_TxsdMtableAttlistAlignmentscope_True
	XsdGoPkgHasAttr_Columnwidth_XsdtString_Auto
	XsdGoPkgHasAttr_Width_XsdtString_Auto
	XsdGoPkgHasAttr_Rowspacing_XsdtString_10Ex
	XsdGoPkgHasAttr_Columnspacing_XsdtString_08Em
	XsdGoPkgHasAttr_Rowlines_XsdtString_None
	XsdGoPkgHasAttr_Columnlines_XsdtString_None
	XsdGoPkgHasAttr_Frame_TxsdMtableAttlistFrame_None
	XsdGoPkgHasAttr_Framespacing_XsdtString_04Em05Ex
	XsdGoPkgHasAttr_Equalrows_XsdtBoolean_False
	XsdGoPkgHasAttr_Equalcolumns_XsdtBoolean_False
	XsdGoPkgHasAttr_Displaystyle_XsdtBoolean_False
	XsdGoPkgHasAttr_Side_TxsdMtableAttlistSide_Right
	XsdGoPkgHasAttr_Minlabelspacing_TlengthWithUnit_08Em
}

type TxsdMaligngroupAttlistGroupalign xsdt.String

//	Since TxsdMaligngroupAttlistGroupalign is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdMaligngroupAttlistGroupalign) SetFromString (s string) { (*xsdt.String)(me).SetFromString(s) }

//	Since TxsdMaligngroupAttlistGroupalign is just a simple String type, this merely returns the current string value.
func (me TxsdMaligngroupAttlistGroupalign) String () string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdMaligngroupAttlistGroupalign's alias type xsdt.String
func (me TxsdMaligngroupAttlistGroupalign) ToXsdtString () xsdt.String { return xsdt.String(me) }

//	Returns true if the value of this enumerated TxsdMaligngroupAttlistGroupalign is "left".
func (me TxsdMaligngroupAttlistGroupalign) IsLeft () bool { return me == "left" }

//	Returns true if the value of this enumerated TxsdMaligngroupAttlistGroupalign is "center".
func (me TxsdMaligngroupAttlistGroupalign) IsCenter () bool { return me == "center" }

//	Returns true if the value of this enumerated TxsdMaligngroupAttlistGroupalign is "right".
func (me TxsdMaligngroupAttlistGroupalign) IsRight () bool { return me == "right" }

//	Returns true if the value of this enumerated TxsdMaligngroupAttlistGroupalign is "decimalpoint".
func (me TxsdMaligngroupAttlistGroupalign) IsDecimalpoint () bool { return me == "decimalpoint" }

type XsdGoPkgHasAttr_Groupalign_TxsdMaligngroupAttlistGroupalign_ struct {
	Groupalign TxsdMaligngroupAttlistGroupalign `xml:"http://www.w3.org/1998/Math/MathML groupalign,attr"`
}

type XsdGoPkgHasAtts_MaligngroupAttlist struct {
	XsdGoPkgHasAtts_CommonAttrib

	XsdGoPkgHasAttr_Groupalign_TxsdMaligngroupAttlistGroupalign_
}

type TxsdMalignmarkAttlistEdge xsdt.String

//	Since TxsdMalignmarkAttlistEdge is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdMalignmarkAttlistEdge) SetFromString (s string) { (*xsdt.String)(me).SetFromString(s) }

//	Since TxsdMalignmarkAttlistEdge is just a simple String type, this merely returns the current string value.
func (me TxsdMalignmarkAttlistEdge) String () string { return xsdt.String(me).String() }

//	This convenience method just performs a simple type conversion to TxsdMalignmarkAttlistEdge's alias type xsdt.String
func (me TxsdMalignmarkAttlistEdge) ToXsdtString () xsdt.String { return xsdt.String(me) }

//	Returns true if the value of this enumerated TxsdMalignmarkAttlistEdge is "left".
func (me TxsdMalignmarkAttlistEdge) IsLeft () bool { return me == "left" }

//	Returns true if the value of this enumerated TxsdMalignmarkAttlistEdge is "right".
func (me TxsdMalignmarkAttlistEdge) IsRight () bool { return me == "right" }

type XsdGoPkgHasAttr_Edge_TxsdMalignmarkAttlistEdge_Left struct {
	Edge TxsdMalignmarkAttlistEdge `xml:"http://www.w3.org/1998/Math/MathML edge,attr"`
}

//	Returns the Default value for Edge -- "left"
func (me *XsdGoPkgHasAttr_Edge_TxsdMalignmarkAttlistEdge_Left) EdgeDefault () TxsdMalignmarkAttlistEdge { return TxsdMalignmarkAttlistEdge("left") }

type XsdGoPkgHasAtts_MalignmarkAttlist struct {
	XsdGoPkgHasAtts_CommonAttrib

	XsdGoPkgHasAttr_Edge_TxsdMalignmarkAttlistEdge_Left
}

type TmtrType struct {
	XsdGoPkgHasGroup_MtrContent

	XsdGoPkgHasAtts_MtrAttlist

}

type TmlabeledtrType struct {
	XsdGoPkgHasGroup_MlabeledtrContent

	XsdGoPkgHasAtts_MlabeledtrAttlist

}

type TmtdType struct {
	XsdGoPkgHasGroup_MtdContent

	XsdGoPkgHasAtts_MtdAttlist

}

type TmtableType struct {
	XsdGoPkgHasGroup_MtableContent

	XsdGoPkgHasAtts_MtableAttlist

}

type TmaligngroupType struct {
	XsdGoPkgHasAtts_MaligngroupAttlist

}

type TmalignmarkType struct {
	XsdGoPkgHasAtts_MalignmarkAttlist

}

type XsdGoPkgHasElems_Mtr struct {
	Mtrs []*TmtrType `xml:"http://www.w3.org/1998/Math/MathML mtr"`
}

type XsdGoPkgHasElem_Mtr struct {
	Mtr *TmtrType `xml:"http://www.w3.org/1998/Math/MathML mtr"`
}

type XsdGoPkgHasElems_Mlabeledtr struct {
	Mlabeledtrs []*TmlabeledtrType `xml:"http://www.w3.org/1998/Math/MathML mlabeledtr"`
}

type XsdGoPkgHasElem_Mlabeledtr struct {
	Mlabeledtr *TmlabeledtrType `xml:"http://www.w3.org/1998/Math/MathML mlabeledtr"`
}

type XsdGoPkgHasElem_Mtd struct {
	Mtd *TmtdType `xml:"http://www.w3.org/1998/Math/MathML mtd"`
}

type XsdGoPkgHasElems_Mtd struct {
	Mtds []*TmtdType `xml:"http://www.w3.org/1998/Math/MathML mtd"`
}

type XsdGoPkgHasElems_Mtable struct {
	Mtables []*TmtableType `xml:"http://www.w3.org/1998/Math/MathML mtable"`
}

type XsdGoPkgHasElem_Mtable struct {
	Mtable *TmtableType `xml:"http://www.w3.org/1998/Math/MathML mtable"`
}

type XsdGoPkgHasElem_Maligngroup struct {
	Maligngroup *TmaligngroupType `xml:"http://www.w3.org/1998/Math/MathML maligngroup"`
}

type XsdGoPkgHasElems_Maligngroup struct {
	Maligngroups []*TmaligngroupType `xml:"http://www.w3.org/1998/Math/MathML maligngroup"`
}

type XsdGoPkgHasElem_Malignmark struct {
	Malignmark *TmalignmarkType `xml:"http://www.w3.org/1998/Math/MathML malignmark"`
}

type XsdGoPkgHasElems_Malignmark struct {
	Malignmarks []*TmalignmarkType `xml:"http://www.w3.org/1998/Math/MathML malignmark"`
}

type XsdGoPkgHasGroup_MtrContent struct {
	XsdGoPkgHasElem_Mtd

}

type XsdGoPkgHasGroup_MlabeledtrContent struct {
	XsdGoPkgHasElem_Mtd

}

type XsdGoPkgHasGroup_MtdContent struct {
	XsdGoPkgHasGroup_PresentationExprClass

}

type XsdGoPkgHasGroup_MtableContent struct {
	XsdGoPkgHasElem_Mtr

	XsdGoPkgHasElem_Mlabeledtr

}

type XsdGoPkgHasGroup_PresentationTableClass struct {
	XsdGoPkgHasElem_Mtable

	XsdGoPkgHasElem_Maligngroup

	XsdGoPkgHasElem_Malignmark

}