import { combineReducers } from 'redux';
import items from './items';
import categories from './categories';

const rootReducer = combineReducers({
  items,
  categories,
});

export default rootReducer;