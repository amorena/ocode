package com.example.fn;

import com.fnproject.fn.api.flow.Flow;
import com.fnproject.fn.api.flow.FlowFuture;
import com.fnproject.fn.api.flow.Flows;
import com.fnproject.fn.api.flow.HttpResponse;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;

class CarsRequest {
    public String search;
}

class Car implements Serializable {
    public String plate;
    public String color;
    public String model;
    public String manufacturer;
    public String year;
    public String description;
    public String condition;

}

class Cars implements Serializable {
    public List<Car> cars;

}

class CriminalResult implements Serializable {
    String stuff;
}

class CarFaxResult implements Serializable {
    String stuff;
}

public class FlowFunction {
    public void handleRequest() throws Exception {
        System.err.println("STARTING FLOW ERR");
        System.out.println("STARTING FLOW OUT");

        Flow fl = Flows.currentFlow();

        FlowFuture<Cars> cars = fl.invokeFunction("./get_cars", new CarsRequest(), Cars.class);

        FlowFuture<Void> allTasks = cars.thenCompose((result) -> {
            List<FlowFuture<Void>> results = new ArrayList<>();

            // iterate over each car and run process_car -> (carfax && criminal_lookup)
            for (Car car : result.cars) {
                FlowFuture<HttpResponse> processCarResult = fl.invokeFunction("./process_car", car);

                FlowFuture<Void> runChecksResult = processCarResult.thenCompose((ignored) -> {
                    FlowFuture<?> carFax = fl.invokeFunction("./carfax", car);
                    FlowFuture<?> criminal = fl.invokeFunction("./criminal_lookup", car);
                    return fl.allOf(carFax,criminal);
                });

                FlowFuture<Void> cardDone = runChecksResult.thenAccept((ignored)->{
                    System.err.println("finished processing " + car.plate );
                });
                results.add(cardDone);

            }

            return fl.allOf(results.toArray(new FlowFuture[results.size()]));
        });

        // Do something with all the results :
        allTasks.thenAccept((ignored) -> {
            System.err.println("Flow done");
        });

    }
}