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

  const handleOpenUrl = () => {
    if(shortURL) {
      window.open(`http://localhost:8101/api/v1/${shortURL}`, '_blank');
    }
  };

  const handleCopyUrl = () => {
    if(shortURL) {
      navigator.clipboard.writeText(`http://localhost:8101/api/v1/${shortURL}`);
    }
  };

  return(
    <div>
      <header>
        <div>
          <h1>Shorten URL</h1>
        </div>
      </header>
      <section id="form-section">
        <div>
          <form onSubmit={handleSubmit}>
            <input type="text" value={url} onChange={(event) => setURL(event.target.value)} id="url" placeholder="Enter URL here" />
            <input type="submit" value="submit" />
          </form>
          {shortURL && (
            <div>
              <div>
                <p>Shorten URL: <a href={`http://localhost:8101/api/v1/${shortURL}`}></a></p>
                <button onClick={handleOpenUrl}>Open</button>
                <button onClick={handleCopyUrl}>Copy</button>
              </div>
            </div>
          )}
        </div>
      </section>
    </div>
  )
}

export default App;
