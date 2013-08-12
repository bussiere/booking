// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/robfig/revel"
	_ "github.com/mattn/go-sqlite3"
	controllers "github.com/robfig/revel/modules/jobs/app/controllers"
	_ "github.com/robfig/revel/modules/jobs/app/jobs"
	controllers0 "github.com/robfig/revel/modules/static/app/controllers"
	_ "github.com/robfig/revel/samples/booking//app"
	_ "github.com/robfig/revel/samples/booking//app/obs"
	_ "github.com/robfig/revel/samples/booking//app/ontrollers"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.Jobs)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Status",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					19: []string{ 
						"entries",
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"github.com/robfig/revel/samples/booking//app/odels.(*Hotel).Validate": { 
			19: "hotel.Name",
			21: "hotel.Address",
			26: "hotel.City",
			32: "hotel.State",
			38: "hotel.Zip",
			44: "hotel.Country",
		},
		"github.com/robfig/revel/samples/booking//app/odels.(*User).Validate": { 
			28: "user.Username",
			36: "user.Name",
		},
		"github.com/robfig/revel/samples/booking//app/odels.Booking.Validate": { 
			34: "booking.User",
			35: "booking.Hotel",
			36: "booking.CheckInDate",
			37: "booking.CheckOutDate",
			39: "booking.CardNumber",
			46: "booking.NameOnCard",
		},
		"github.com/robfig/revel/samples/booking//app/odels.ValidatePassword": { 
			44: "password",
		},
		"github.com/robfig/revel/samples/booking//app/ontrollers.Application.SaveUser": { 
			55: "verifyPassword",
			56: "verifyPassword",
		},
		"github.com/robfig/revel/samples/booking//app/ontrollers.Hotels.SaveSettings": { 
			98: "verifyPassword",
			100: "verifyPassword",
		},
	}
	revel.TestSuites = []interface{}{ 
	}

	revel.Run(*port)
}
