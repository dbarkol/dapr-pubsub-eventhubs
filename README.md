# Publish-Subscribe with Event Hubs on Dapr

This solution demonstrates how to consume messages from a message broker with the [Dapr](https://github.com/dapr) pub-sub component using Node.js, Go and C#.

The accompanying blog post goes through the configuration steps and considerations when choosing [Azure Event Hubs](https://aka.ms/azureeventhubs) as the message broker. However, the code for consuming the messages does not change if a different message broker is selected, only the component YAML file.

Blog post: [https://madeofstrings.com/2020/07/05/pub-sub-with-dapr-and-azure-event-hubs/](https://madeofstrings.com/2020/07/05/pub-sub-with-dapr-and-azure-event-hubs/)

At a high-level, the physical architecture for this solution looks like:

![Physical architecture](/images/dapr-eventhubs-physical-architecture.png)

## Setup with Azure Cloud Shell for Azure Event Hubs

### Initialize variables

``` bash
# resource group name and location
rgname=zohan-dapr-demo
location=westus2

# event hubs namespace and event hub (topic)
ehnamespace=<unique-event-hubs-namespace-name>
ehname=songs

# authorization rule for shared access policy
authorizationrule=authorizationpolicy

# consumer group for each subscriber
consumergroup1=csharp-subscriber-app
consumergroup2=node-subscriber-app
consumergroup3=go-subscriber-app

# storage account name
storageaccount=<unique-storage-account-name>
```

### Create Event Hubs and Event Hub (topic)

``` bash
# create event hubs namespace
az eventhubs namespace create --name $ehnamespace -g $rgname -l $location --sku Standard

# create event hub (topic)
az eventhubs eventhub create -g $rgname --namespace-name $ehnamespace --name $ehname
```

### Create consumer groups, one for each subscriber

``` bash
# csharp subscriber
az eventhubs eventhub consumer-group create --resource-group $rgname --namespace-name $ehnamespace --eventhub-name $ehname --name $consumergroup1

# node subscriber
az eventhubs eventhub consumer-group create --resource-group $rgname --namespace-name $ehnamespace --eventhub-name $ehname --name $consumergroup2

# go subscriber
az eventhubs eventhub consumer-group create --resource-group $rgname --namespace-name $ehnamespace --eventhub-name $ehname --name $consumergroup3
```

### Create a shared access policy

``` bash
# create shared access policy with send and listen rights
az eventhubs eventhub authorization-rule create --eventhub-name $ehname --name $authorizationrule --namespace-name $ehnamespace -g $rgname --rights Send Listen

# query the primary connection string
az eventhubs eventhub authorization-rule keys list --resource-group $rgname --namespace-name $ehnamespace --eventhub-name $ehname --name $authorizationrule --query "primaryConnectionString"
```

### Create a storage account

``` bash
# create storage account
az storage account create --name $storageaccount --location $location --resource-group $rgname --sku Standard_LRS --kind BlobStorage --access-tier Hot
```

## Event Hubs pub-sub component

Important component details:

- A different component must be used for each topic subscription.
- The storage container name must match the name of consumer group for the subscription. For example, the storage account name for the Node.js application should be `node-subscriber-app`.

Here is an example of the component configuration for the Node application:

``` YAML
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: messagebus-node
spec:
  type: pubsub.azure.eventhubs
  metadata:
    - name: connectionString
      value: Endpoint=sb://<namespace-name>.servicebus.windows.net/;SharedAccessKeyName=<policy-name>;SharedAccessKey=<key>;EntityPath=<event-hub-name>
    - name: storageAccountName
      value: <storage-account-name>
    - name: storageAccountKey
      value: <storage-account-key>
    - name: storageContainerName
      value: node-subscriber-app
```
