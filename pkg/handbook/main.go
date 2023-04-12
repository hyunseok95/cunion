package handbook

import (
	// controlstructures "github.com/hyunseok95/cunion/pkg/handbook/control-structures"
	Functions "github.com/hyunseok95/cunion/pkg/handbook/functions"
)

func Run() error {
	// if err := controlstructures.IfFunc(); err != nil {
	// 	return err
	// }

	// if err := controlstructures.ForFunc(); err != nil {
	// 	return err
	// }

	// if err := controlstructures.SwitchFunc(args); err != nil {
	// 	return err
	// }

	if err := Functions.Functions(); err != nil {
		return err
	}

	return nil
}
