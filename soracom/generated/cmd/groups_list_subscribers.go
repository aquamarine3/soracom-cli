package cmd

import (

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var GroupsListSubscribersCmdGroupId string

var GroupsListSubscribersCmdLastEvaluatedKey string


var GroupsListSubscribersCmdLimit int64





func init() {
  GroupsListSubscribersCmd.Flags().StringVar(&GroupsListSubscribersCmdGroupId, "group-id", "", TR("groups.list_subscribers_in_group.get.parameters.group_id.description"))

  GroupsListSubscribersCmd.Flags().StringVar(&GroupsListSubscribersCmdLastEvaluatedKey, "last-evaluated-key", "", TR("groups.list_subscribers_in_group.get.parameters.last_evaluated_key.description"))

  GroupsListSubscribersCmd.Flags().Int64Var(&GroupsListSubscribersCmdLimit, "limit", 0, TR("groups.list_subscribers_in_group.get.parameters.limit.description"))




  GroupsCmd.AddCommand(GroupsListSubscribersCmd)
}

var GroupsListSubscribersCmd = &cobra.Command{
  Use: "list-subscribers",
  Short: TR("groups.list_subscribers_in_group.get.summary"),
  Long: TR(`groups.list_subscribers_in_group.get.description`),
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
    
    param, err := collectGroupsListSubscribersCmdParams()
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

func collectGroupsListSubscribersCmdParams() (*apiParams, error) {
  

  return &apiParams{
    method: "GET",
    path: buildPathForGroupsListSubscribersCmd("/groups/{group_id}/subscribers"),
    query: buildQueryForGroupsListSubscribersCmd(),
    
    
  }, nil
}

func buildPathForGroupsListSubscribersCmd(path string) string {
  
  
  path = strings.Replace(path, "{" + "group_id" + "}", GroupsListSubscribersCmdGroupId, -1)
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForGroupsListSubscribersCmd() string {
  result := []string{}
  
  
  
  
  if GroupsListSubscribersCmdLastEvaluatedKey != "" {
    result = append(result, sprintf("%s=%s", "last_evaluated_key", GroupsListSubscribersCmdLastEvaluatedKey))
  }
  
  

  
  
  if GroupsListSubscribersCmdLimit != 0 {
    result = append(result, sprintf("%s=%d", "limit", GroupsListSubscribersCmdLimit))
  }
  
  

  

  

  return strings.Join(result, "&")
}


