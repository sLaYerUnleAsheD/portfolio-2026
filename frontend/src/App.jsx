import Layout from './components/Layout';
import Hero from './components/Hero';
import Skills from './components/Skills';
import ArtOfTheDay from './components/ArtOfTheDay';
import MusicBooth from './components/MusicBooth';
import Footer from './components/Footer';
import HiddenCat from './components/HiddenCat';

/**
 * App — Root component that assembles all portfolio sections
 * inside the Layout shell. Each section is a self-contained module.
 */
export default function App() {
  return (
    <Layout>
      <Hero />
      <Skills />
      <ArtOfTheDay />
      <MusicBooth />
      <Footer />
      <HiddenCat />
    </Layout>
  );
}
