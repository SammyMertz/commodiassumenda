import _ from 'lodash'
import * as BRIDGEABLE from '@constants/tokens/bridgeable'

export const findTokenByRouteSymbol = (routeSymbol: string) => {
  return _.find(BRIDGEABLE, (token) => token.routeSymbol === routeSymbol)
}
