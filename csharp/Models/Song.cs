using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using CloudNative.CloudEvents;
using Newtonsoft.Json;

namespace SongRequests.Models
{
    public class Song
    {
        [JsonProperty("id")]
        public int Id { get; set; }

        [JsonProperty("artist")]
        public string Artist { get; set; }

        [JsonProperty("name")]
        public string Name { get; set; }
    }
}