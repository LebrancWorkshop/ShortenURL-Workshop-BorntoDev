import React, { useState } from 'react';
import axios from 'axios';

const App: React.FC = () => {
  const [url, setURL] = useState<string>();
  const [shortURL, setShortURL] = useState<string>();

  const handleSubmit = async(event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    try {
      const response = await axios.post('/api/v1/url', { url });
      const { short_url } = response.data;
      setShortURL(short_url); 
    } catch (error) {
      console.error(error);
    }

  }

  return(
    <div>
      <header>
        <div>
          <h1>Shorten URL</h1>
        </div>
      </header>
      <section id="form-section">
        <div>
          <form action="POST">
            <input type="text" name="url" id="url" placeholder="Enter URL here" />
            <input type="submit" value="submit" />
          </form>
        </div>
      </section>
    </div>
  )
}

export default App;
