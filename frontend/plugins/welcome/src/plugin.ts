import { createPlugin } from '@backstage/core';
import Disease from './components/Disease'
import Login from './components/Login'
import Table from './components/Table'

const IMG1 = './assets/Login.jpg'


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/Disease', Disease);
    router.registerRoute('/Table', Table);
  },
});
