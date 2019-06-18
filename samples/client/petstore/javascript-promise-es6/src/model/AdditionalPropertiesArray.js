/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The AdditionalPropertiesArray model module.
 * @module model/AdditionalPropertiesArray
 * @version 1.0.0
 */
class AdditionalPropertiesArray {
    /**
     * Constructs a new <code>AdditionalPropertiesArray</code>.
     * @alias module:model/AdditionalPropertiesArray
     * @extends Object
     */
    constructor() { 
        
        AdditionalPropertiesArray.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>AdditionalPropertiesArray</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AdditionalPropertiesArray} obj Optional instance to populate.
     * @return {module:model/AdditionalPropertiesArray} The populated <code>AdditionalPropertiesArray</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new AdditionalPropertiesArray();

            ApiClient.constructFromObject(data, obj, 'Array');
            

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} name
 */
AdditionalPropertiesArray.prototype['name'] = undefined;






export default AdditionalPropertiesArray;

