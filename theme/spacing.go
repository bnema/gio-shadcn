package theme

import "gioui.org/unit"

// SpacingScale defines the spacing system for consistent layout spacing.
// The spacing scale follows a 4px base unit system, providing harmonious.
// spacing options that work well across different screen sizes and components.
// All values are in density-independent pixels (dp) for proper scaling.
//
// Usage example:.
//
//	layout.Inset{
//		Top:    th.Spacing.Space4,  // 16dp
//		Bottom: th.Spacing.Space8,  // 32dp
//	}.Layout(gtx, content)
type SpacingScale struct {
	// Base spacing units (in dp)
	Space0  unit.Dp // 0px
	Space1  unit.Dp // 4px
	Space2  unit.Dp // 8px
	Space3  unit.Dp // 12px
	Space4  unit.Dp // 16px
	Space5  unit.Dp // 20px
	Space6  unit.Dp // 24px
	Space7  unit.Dp // 28px
	Space8  unit.Dp // 32px
	Space9  unit.Dp // 36px
	Space10 unit.Dp // 40px
	Space11 unit.Dp // 44px
	Space12 unit.Dp // 48px
	Space14 unit.Dp // 56px
	Space16 unit.Dp // 64px
	Space20 unit.Dp // 80px
	Space24 unit.Dp // 96px
	Space28 unit.Dp // 112px
	Space32 unit.Dp // 128px
	Space36 unit.Dp // 144px
	Space40 unit.Dp // 160px
	Space44 unit.Dp // 176px
	Space48 unit.Dp // 192px
	Space52 unit.Dp // 208px
	Space56 unit.Dp // 224px
	Space60 unit.Dp // 240px
	Space64 unit.Dp // 256px
	Space72 unit.Dp // 288px
	Space80 unit.Dp // 320px
	Space96 unit.Dp // 384px
}

// RadiusScale defines the border radius system for consistent rounded corners.
// The radius scale provides a range of border radius values from sharp corners.
// to fully rounded elements. Values are in density-independent pixels (dp).
//
// Usage example:.
//
//	// Apply medium radius to a component
//	paint.FillShape(gtx.Ops, color, clip.RRect{
//		Rect: bounds,
//		NE:   int(gtx.Metric.Dp(th.Radius.RadiusMD)),
//		NW:   int(gtx.Metric.Dp(th.Radius.RadiusMD)),
//		SE:   int(gtx.Metric.Dp(th.Radius.RadiusMD)),
//		SW:   int(gtx.Metric.Dp(th.Radius.RadiusMD)),
//	}.Op(gtx.Ops))
type RadiusScale struct {
	RadiusNone unit.Dp // 0px
	RadiusSM   unit.Dp // 2px
	RadiusBase unit.Dp // 4px
	RadiusMD   unit.Dp // 6px
	RadiusLG   unit.Dp // 8px
	RadiusXL   unit.Dp // 12px
	Radius2XL  unit.Dp // 16px
	Radius3XL  unit.Dp // 24px
	RadiusFull unit.Dp // 9999px (effectively full rounding)
}

// DefaultSpacing returns the default spacing configuration.
// The spacing scale uses a 4dp base unit with consistent progression,.
// providing harmonious spacing options for layouts and components.
// Values range from 0dp to 384dp to cover all common spacing needs.
func DefaultSpacing() SpacingScale {
	return SpacingScale{
		Space0:  unit.Dp(0),
		Space1:  unit.Dp(4),
		Space2:  unit.Dp(8),
		Space3:  unit.Dp(12),
		Space4:  unit.Dp(16),
		Space5:  unit.Dp(20),
		Space6:  unit.Dp(24),
		Space7:  unit.Dp(28),
		Space8:  unit.Dp(32),
		Space9:  unit.Dp(36),
		Space10: unit.Dp(40),
		Space11: unit.Dp(44),
		Space12: unit.Dp(48),
		Space14: unit.Dp(56),
		Space16: unit.Dp(64),
		Space20: unit.Dp(80),
		Space24: unit.Dp(96),
		Space28: unit.Dp(112),
		Space32: unit.Dp(128),
		Space36: unit.Dp(144),
		Space40: unit.Dp(160),
		Space44: unit.Dp(176),
		Space48: unit.Dp(192),
		Space52: unit.Dp(208),
		Space56: unit.Dp(224),
		Space60: unit.Dp(240),
		Space64: unit.Dp(256),
		Space72: unit.Dp(288),
		Space80: unit.Dp(320),
		Space96: unit.Dp(384),
	}
}

// DefaultRadius returns the default border radius configuration.
// The radius scale provides a range from sharp corners (0dp) to fully.
// rounded elements (9999dp). Values are chosen to provide good visual
// hierarchy and work well with the overall design system.
func DefaultRadius() RadiusScale {
	return RadiusScale{
		RadiusNone: unit.Dp(0),
		RadiusSM:   unit.Dp(2),
		RadiusBase: unit.Dp(4),
		RadiusMD:   unit.Dp(6),
		RadiusLG:   unit.Dp(8),
		RadiusXL:   unit.Dp(12),
		Radius2XL:  unit.Dp(16),
		Radius3XL:  unit.Dp(24),
		RadiusFull: unit.Dp(9999),
	}
}
