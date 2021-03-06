package main

import "github.com/elamre/vcd"

func main() {
	writer, e := vcd.New("example", "10ps")
	if e != nil {
		panic(e)
	}
	defer writer.Close()
	writer.SetVersion("1.0.0").SetComment("Example for VCD GO")
	_, e = writer.RegisterVariables("example.logic",
		vcd.NewVariable("miso", "wire", 8),
		vcd.NewVariable("mosi", "wire", 8),
		vcd.NewVariable("cs", "wire", 1),
	)
	if e != nil {
		panic(e)
	}
	_, e = writer.RegisterVariables("example",
		vcd.NewVariable("command", "string", 1),
	)
	_, e = writer.RegisterVariables("example.real",
		vcd.NewVariable("analogue", "real", 1),
	)
	if e != nil {
		panic(e)
	}
	_ = writer.SetValue(0, "1", "cs")
	_ = writer.SetValue(0, "", "command")
	_ = writer.SetValue(0, "z", "mosi")
	_ = writer.SetValue(0, "z", "miso")
	_ = writer.SetValue(0, "1.1", "analogue")
	_ = writer.SetValue(100, "0", "cs")
	_ = writer.SetValue(100, "String command", "command")
	_ = writer.SetValue(100, "80", "mosi")
	_ = writer.SetValue(100, "0", "miso")
	_ = writer.SetValue(100, "1.31", "analogue")
	_ = writer.SetValue(200, "48", "mosi")
	_ = writer.SetValue(200, "43", "miso")
	_ = writer.SetValue(100, "7.31", "analogue")
	_ = writer.SetValue(300, "90", "mosi")
	_ = writer.SetValue(300, "10", "miso")
	_ = writer.SetValue(300, "5.1", "analogue")
	_ = writer.SetValue(400, "100", "mosi")
	_ = writer.SetValue(400, "90", "miso")
	_ = writer.SetValue(400, "0.1", "analogue")
	_ = writer.SetValue(500, "1", "cs")
	_ = writer.SetValue(500, "z", "miso")
	_ = writer.SetValue(500, "z", "mosi")
	_ = writer.SetValue(500, "", "command")
	_ = writer.SetValue(500, "0", "analogue")
	writer.SetTimestamp(600)
	gtkw := vcd.Gtkw("example")
	defer gtkw.Close()
	gtkw.Group("SPI", true,
		vcd.Trace("example.logic.mosi[7:0]", "mosi", "rjustify", "hex"),
		vcd.Trace("example.logic.miso[7:0]", "miso", "rjustify", "hex"),
		vcd.Trace("example.logic.cs", "cs", "bin"),
	)
	gtkw.Group("analogue", false,
		vcd.Trace("example.real.analogue", "analogue", "analog_interpolated", "analog_fullscale"),
	)
	gtkw.Trace(vcd.Trace("example.command", "command"))
	gtkw.SetDumpfile("example.vcd")
}
