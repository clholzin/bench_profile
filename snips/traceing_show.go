// Set up a trace context and task
ctx, task := trace.NewTask(ctx, "Foo http route")
// Log with context and uuid
trace.Log(ctx,"Execute Foo route work flow",<uuid>)
trace.WithRegion(ctx, "steamMilk", steamMilk)
// Pass context along the flow into call stack
workflowFunc(ctx context.Context,id string){ 
	trace.log(ctx,"Workflow execution stage",id)
	//work being done
}
