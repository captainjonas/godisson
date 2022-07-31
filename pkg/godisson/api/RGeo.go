package api

/**
* Geospatial items holder
 */
type RGeo interface {

	/**
	* Adds geospatial member
	 */
	add(longitude, latitude float64, member interface{}) (int64, error)

	/**
	* Adds geospatial member only if has not been added before
	 */
	tryAdd(longitude, latitude float64, member interface{}) (bool, error)

	/**
	* Returns 11 characters long Geohash string mapped by defined member.
	 */
	hash(members []interface{}) (map[interface{}]string, error)
}
