import * as pulumi from "@pulumi/pulumi";
import * as ovh from "@ovhcloud/pulumi-ovh";

// Get configuration values or use defaults
const config = new pulumi.Config();
const ovhServiceName = config.require("ovhServiceName");
const ovhRegion = config.get("ovhRegion") || "GRA";
const ovhPlan = config.get("ovhPlan") || "SMALL";

// Fetch the capabilities for the specified plan and region
const capabilities = ovh.cloudproject.getCapabilitiesContainerFilter({
    serviceName: ovhServiceName,
    region: ovhRegion,
    planName: ovhPlan,
}); 

//Display the registry capability ID
capabilities.then(cap => {
    console.log("Plan ID:", cap.id);
});

// Create a new container registry using the fetched capabilities
const myRegistry = new ovh.cloudproject.ContainerRegistry("my_registry", {
    serviceName: ovhServiceName,
    planId: capabilities.then(cap => cap.id),
    region: ovhRegion,
    name: "mydockerregistry",
}); 

// Display the registry information after creation
myRegistry.id.apply(id => console.log("Registry ID:", id));
myRegistry.name.apply(name => console.log("Registry Name:", name));
myRegistry.region.apply(region => console.log("Registry Region:", region));
myRegistry.planId.apply(planId => console.log("Registry Plan ID:", planId));
myRegistry.url.apply(url => console.log("Registry URL:", url));

// Create a registry user
const registryUser = new ovh.cloudproject.ContainerRegistryUser("registry_user", {
    serviceName: ovhServiceName,
    registryId: myRegistry.id,
    email: "myuser@ovh.com",
    login: "my_user",
});

// Display the registry user information after creation
registryUser.login.apply(login => console.log("Registry User Login:", login));
registryUser.email.apply(email => console.log("Registry User Email:", email));
registryUser.password.apply(password => console.log("Registry User Password:", password));
