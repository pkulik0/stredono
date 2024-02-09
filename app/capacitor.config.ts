import { CapacitorConfig } from '@capacitor/cli';

const config: CapacitorConfig = {
  appId: 'com.piotrkulik.stredono',
  appName: 'stredono',
  webDir: 'build',
  server: {
    androidScheme: 'https'
  }
};

export default config;
