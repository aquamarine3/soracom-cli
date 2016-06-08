package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var StatsAirGetCmdImsi string

var StatsAirGetCmdPeriod string


var StatsAirGetCmdFrom int64

var StatsAirGetCmdTo int64





func init() {
  StatsAirGetCmd.Flags().StringVar(&StatsAirGetCmdImsi, "imsi", "", TR("stats.get_air_stats.get.parameters.imsi.description"))

  StatsAirGetCmd.Flags().StringVar(&StatsAirGetCmdPeriod, "period", "", TR("stats.get_air_stats.get.parameters.period.description"))

  StatsAirGetCmd.Flags().Int64Var(&StatsAirGetCmdFrom, "from", 0, TR("stats.get_air_stats.get.parameters.from.description"))

  StatsAirGetCmd.Flags().Int64Var(&StatsAirGetCmdTo, "to", 0, TR("stats.get_air_stats.get.parameters.to.description"))




  StatsAirCmd.AddCommand(StatsAirGetCmd)
}

var StatsAirGetCmd = &cobra.Command{
  Use: "get",
  Short: TR("stats.get_air_stats.get.summary"),
  Long: TR(`stats.get_air_stats.get.description`),
  RunE: func(cmd *cobra.Command, args []string) error {
    opt := &apiClientOptions{
      Endpoint: getSpecifiedEndpoint(),
      BasePath: "/v1",
      Language: getSelectedLanguage(),
    }

    ac := newAPIClient(opt)
    if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
      ac.SetVerbose(true)
    }

    
    err := authHelper(ac, cmd, args)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }
    
    param, err := collectStatsAirGetCmdParams()
    if err != nil {
      return err
    }

    result, err := ac.callAPI(param)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }

    if result != "" {
      return prettyPrintStringAsJSON(result)
    } else {
      return nil
    }
  },
}

func collectStatsAirGetCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForStatsAirGetCmd("/stats/air/subscribers/{imsi}"),
    query: buildQueryForStatsAirGetCmd(),
    
    
  }, nil
}

func buildPathForStatsAirGetCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "imsi" + "}", StatsAirGetCmdImsi, -1)
  
  
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForStatsAirGetCmd() string {
  result := []string{}
  
  
  
  
  if StatsAirGetCmdPeriod != "" {
    result = append(result, sprintf("%s=%s", "period", StatsAirGetCmdPeriod))
  }
  
  

  
  
  if StatsAirGetCmdFrom != 0 {
    result = append(result, sprintf("%s=%d", "from", StatsAirGetCmdFrom))
  }
  
  
  
  if StatsAirGetCmdTo != 0 {
    result = append(result, sprintf("%s=%d", "to", StatsAirGetCmdTo))
  }
  
  

  

  

  return strings.Join(result, "&")
}


