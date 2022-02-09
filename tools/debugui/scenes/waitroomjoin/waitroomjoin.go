package waitroomjoin

func Show() {
	// if client := networking.UseNetworking().Clients().Base().GetClient(); client != nil {
	// 	// w := world.UseWorld()
	// 	// w.DebugInit()
	// 	worldMess := w.ToAPIMessage()
	// 	mapMess := w.GetWorldMap().ToAPIMessage()
	// 	pcMess := w.GetPC().ToAPIMessage()

	// 	if _, err := client.InsertOrUpdateWorld(context.Background(), worldMess, grpc.EmptyCallOption{}); err != nil {
	// 		notifications.PopUp.WriteError(err.Error())
	// 		logrus.Error(err)
	// 		return
	// 	}

	// 	if _, err := client.InsertOrUpdateMap(context.Background(), mapMess, grpc.EmptyCallOption{}); err != nil {
	// 		notifications.PopUp.WriteError(err.Error())
	// 		logrus.Error(err)
	// 		return
	// 	}

	// 	if _, err := client.InsertOrUpdatePC(context.Background(), pcMess, grpc.EmptyCallOption{}); err != nil {
	// 		notifications.PopUp.WriteError(err.Error())
	// 		logrus.Error(err)
	// 		return
	// 	}
	// 	// middlewares.UseMiddlewares().UI().UseAfter(func() {
	// 	// 	statemachine.UseStateMachine().UI().SetState(statemachine.UI_GAME)
	// 	// })

	// 	// statemachine.UseStateMachine().Input().SetState(statemachine.INPUT_GAME)
	// }
}
