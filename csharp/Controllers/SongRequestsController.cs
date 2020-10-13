using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using CloudNative.CloudEvents;
using SongRequests.Models;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;

namespace SongRequests.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class RequestController : ControllerBase
    {
        private readonly ILogger<RequestController> _logger;

        public RequestController(ILogger<RequestController> logger)
        {
            _logger = logger;
        }

        [HttpGet("/dapr/subscribe")]
        public ActionResult<IEnumerable<string>> Get()
        {
            // Initialize an array of topic subscriptions. Each subscription
            // contains the name of the topic and the route.
            var topics = new [] {new { pubsubname="messagebus-csharp", topic = "songs", route = "playlist"}};
                                
            return new OkObjectResult(topics);       
        }

        [HttpPost("/playlist")]
        public async Task<IActionResult> NewSong(CloudEvent cloudEvent)
        {
            // The message is wrapped in a cloud event envelope. Which means that 
            // the domain-specific information (the song) is in the Data object.
            var song = JsonConvert.DeserializeObject<Song>(cloudEvent.Data.ToString());
            _logger.LogInformation($"New song request: {song.Artist} - {song.Name}");

            return new OkResult();
        }
    }
}