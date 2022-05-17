// component
import Iconify from '../../components/Iconify';
import { Item } from '../../components/NavSection';

// ----------------------------------------------------------------------

const getIcon = (name: any) => {
  return <Iconify icon={name} width={22} height={22} />;
};

const navConfig: Item[] = [
  {
    title: 'dashboard',
    path: '',
    icon: getIcon('eva:pie-chart-2-fill'),
  },
  {
    title: 'subscribe',
    path: 'subscribe',
    icon: getIcon('gridicons:reader-following'),
  },
];

export default navConfig;
