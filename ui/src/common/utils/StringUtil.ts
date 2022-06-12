export default class StringUtil {
    static hasText(value: string|number|undefined): boolean {
        if (!value) {
            return false
        }
        if (typeof value == "string") {
            return value.trim() !== "";
        }else {
            return true
        }
    }

    static isBlank(value: string|number): boolean {
        return !this.hasText(value);
    }

    static contains(value: string, containVal: string): boolean {
        return value.indexOf(containVal) > -1;
    }
}