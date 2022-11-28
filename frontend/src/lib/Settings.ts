
export default class Settings {
    CACert:         string;
    CAKey:          string;
    ProxyPort:      number;
    ProxyHost:      string;
    LogLevel: number;
    Theme:       string;
    constructor() {
        this.CACert = "";
        this.CAKey = "";
        this.ProxyPort = 0;
        this.ProxyHost = "";
        this.LogLevel = 0;
        this.Theme = "";
    }
}