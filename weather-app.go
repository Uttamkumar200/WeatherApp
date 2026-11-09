import {useState} from "react";
import "./App.css";

function App(){
  const [city, setCity] = useState("");

  const [weather, setWeather]= useState(null);
 
  // Loading State
  const [loading, setLoading] = useState(false);

  // Error State
  const [error, setError] = useState("");

  //search function

  const searchWeather = async () => {
    if(city===""){
      setError("please enter a city name");
      return;
    }
    try{
      setLoading(true);
      setError("");
      
        
      
      
      const locationResponse = await fetch( 
        `https://geocoding-api.open-meteo.com/v1/search?name=${city}&count=1&language=en&format=json`

    );
  
      const locationData = await locationResponse.json();
      if(!locationData.results){
        setError("city not found");
        setLoading(false);
        return;
      }
      const latitude = locationData.results[0].latitude;
      const longitude = locationData.results[0].longitude;
      const cityname = locationData.results[0].name;
      const country = locationData.results[0].country;
      const weatherResponse = await fetch(
  `https://api.open-meteo.com/v1/forecast?latitude=${latitude}&longitude=${longitude}&daily=weather_code,temperature_2m_max,temperature_2m_min,precipitation_probability_max&timezone=auto&forecast_days=5`
);
      const weatherData = await weatherResponse.json();
      setWeather({
        city: cityname,
        country: country,
        data:weatherData.daily,
      });
      setLoading(false);
    }
    catch (error) {
  console.error("Weather Error:", error);
  setError(error.message);
}
  };

  if (code <= 48) return "🌫️ Foggy";
 const getWeatherText = (code) => {
  
  if (code === 0) return "☀️ Clear Sky";
  if (code <= 3) return "🌤️ Partly Cloudy";
  if (code <= 48) return "🌫️ Foggy";
  if (code <= 57) return "🌦️ Drizzle";
  if (code <= 67) return "🌧️ Rain";
  if (code <= 77) return "❄️ Snow";
  if (code <= 82) return "🌧️ Rain Shower";
  if (code <= 86) return "🌨️ Snow Shower";
  if (code <= 95) return "⛈️ Thunderstorm";
  return "❓ Unknown";
};

  return(
    <div className="app">
      <div className="weather-box">
        <h1>Weather App</h1>
        <p>5 Days Weather Forecast</p>
        <div className="Search-box">
          <input type="text" placeholder="Enter city name" value={city} onChange={(e)=> setCity(e.target.value)}></input>
          <button onClick={searchWeather}>Search</button>
        </div>
        {
          error && (
            <p className="error">{error}</p>
            
          )
        }
        {loading && (
          <p className="loading">Loading weather...</p>
        )}

        { weather && !loading && (
          <div>

            <h2>
              {weather.city}, {weather.country}
            </h2>

            <div className="forecast">
              {weather.data.time.map((date,index)=> (
                <div className="weather-card" key={date}>
                  <h3>{new Date(date).toDateString("en-US", {weekday: "short",})}

                  </h3>
                  <p className="date">{date}</p>
                  <div className="weather-icon">
                    {getWeatherText(weather.data.weather_code[index]).split(" ")[0]}


                  </div>
                  <p>{getWeatherText(weather.data.weather_code[index]).substring(2)}</p>
                  <h2>{weather.data.temperature_2m_max[index]}°C</h2>
                  <p>Min: {weather.data.temperature_2m_min[index]}°C</p>
                  <p>Rain: {" "} {weather.data.precipitation_probability_max[index]}%</p>

                   


                </div>
              ))}
            </div>



          </div>
        )}


             

      </div>
    </div>
  );




}

export default App;


