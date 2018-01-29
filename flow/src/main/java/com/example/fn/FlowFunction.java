package com.example.fn;

import com.fnproject.fn.api.flow.Flow;
import com.fnproject.fn.api.flow.Flows;
import com.fnproject.fn.api.flow.FlowFuture;
import com.fnproject.fn.api.RuntimeContext;
import com.fnproject.fn.api.FnConfiguration;

import com.fnproject.fn.api.flow.HttpMethod;
import com.fnproject.fn.api.flow.HttpResponse;

import java.util.List;
import java.util.stream.Collectors;

import static com.fnproject.fn.api.flow.Flows.currentFlow;


public class FlowFunction {
    public void handleRequest() throws Exception {
        System.err.println("STARTING FLOW ERR");
        System.out.println("STARTING FLOW OUT");

        Flow fl = Flows.currentFlow();

        FlowFuture<HttpResponse> p1 = fl.invokeFunction("devweek/process_car", HttpMethod.POST);
        FlowFuture<HttpResponse> p2 = fl.invokeFunction("devweek/process_car", HttpMethod.POST);
        FlowFuture<HttpResponse> p3 = fl.invokeFunction("devweek/process_car", HttpMethod.POST);


        // CURRENTLY DOING THIS
        FlowFuture<HttpResponse> c1 = p1.thenCompose( msg -> {
            return fl.invokeFunction("devweek/carfax", HttpMethod.POST);
        });
        FlowFuture<HttpResponse> c2 = p1.thenCompose( msg -> {
            return fl.invokeFunction("devweek/criminal_lookup", HttpMethod.POST);
        });
        FlowFuture<HttpResponse> c3 = p2.thenCompose( msg -> {
            return fl.invokeFunction("devweek/carfax", HttpMethod.POST);
        });
        FlowFuture<HttpResponse> c4 = p2.thenCompose( msg -> {
            return fl.invokeFunction("devweek/criminal_lookup", HttpMethod.POST);
        });

        FlowFuture<HttpResponse> c5 = p3.thenCompose( msg -> {
            return fl.invokeFunction("devweek/carfax", HttpMethod.POST);
        });
        FlowFuture<HttpResponse> c6 = p3.thenCompose( msg -> {
            return fl.invokeFunction("devweek/criminal_lookup", HttpMethod.POST);
        });

        // WANT SOMETHING MORE LIKE THIS:
        // Run process_car 5 times, for each one, run in parallel carfax and criminal_lookup

        // FlowFuture<HttpResponse> c1 = p1.thenCompose( msg -> {
        //     return currentFlow()
        //         .allOf(fl.invokeFunction("devweek/carfax", HttpMethod.POST),
        //                 fl.invokeFunction("devweek/criminal_lookup", HttpMethod.POST));
        // });




        System.out.println("ENDING FLOW");
    }
}